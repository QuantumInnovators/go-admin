import pandas as pd
import pymysql

# 配置MySQL连接信息
host = "localhost"
port = 3306
user = "root"
password = "raspberrypi"
database = "aquatic"

files = ["2023年第一批底栖数据库及与NCBI数据库对比1024.xlsx"]
# 读取Excel文件
excel_file = "/Users/zhouyaoxu/Desktop/参考设计/本地数据集/"

# 连接MySQL数据库
conn = pymysql.connect(
    host=host, port=port, user=user, password=password, database=database
)
cursor = conn.cursor()
# 界门纲目科属种
for file in files:
    df = pd.read_excel(excel_file + file)
    for index, row in df.iterrows():
        sequence = row[5]  # 序列
        genus_name = row[4]  # 属名or种名
        genus_cn_name = row[3]  # 中文属名or种名
        # 引物名称
        primer_name = row[2]
        typ_name = row[6]

        # 从种表获取name=genus_name或者genus_cn_name=desc的id
        # 如果没有，从属表获取name=genus_name或者genus_cn_name=desc的id
        sql = f"SELECT * FROM species WHERE `name`='{genus_name}' OR `name_cn`='{genus_cn_name}'"
        cursor.execute(sql)
        result = cursor.fetchall()
        is_species = True
        if len(result) == 0:
            is_species = False
            sql = f"SELECT * FROM genus WHERE `name`='{genus_name}' OR `name_cn`='{genus_cn_name}'"
            cursor.execute(sql)
            result = cursor.fetchall()
        if len(result) == 0:
            print("Error: 未找到对应的属或种")
            continue
        ret = result[0]
        id = ret[0]
        parent_id = ret[3]
        if is_species:
            # 从属表获取parent_id
            sql = f"SELECT * FROM genus WHERE id='{parent_id}'"
            cursor.execute(sql)
            result = cursor.fetchone()
            if result is None:
                print("Error: 未找到对应的属")
                continue
            parent_id = result[3]

        family_id = parent_id
        # 从科表获取parent_id
        sql = f"SELECT * FROM family WHERE id='{family_id}'"
        cursor.execute(sql)
        result = cursor.fetchone()
        order_id = result[3]
        # 从目表获取parent_id
        sql = f"SELECT * FROM `order` WHERE id='{order_id}'"
        cursor.execute(sql)
        result = cursor.fetchone()
        class_id = result[3]
        # 从纲表获取parent_id
        sql = f"SELECT * FROM class WHERE id='{class_id}'"
        cursor.execute(sql)
        result = cursor.fetchone()
        phylum_id = result[3]

        # 规则：例如底栖动物，门（2位数字），目（2位数字），科（3位数字）.第几条
        # DQDW0101001.1
        # DQDW0101002.1
        sequence_id = f"DQDW{phylum_id:02d}{order_id:02d}{family_id:03d}.{index}"
        sequence_desc = genus_name + " " + primer_name
        # 插入序列表
        sql = f"INSERT INTO `sequence_local` (`sequence_id`, `name`, `name_zh`, `type`, `primer_name`, `sequence_description`, `sequence`) VALUES ('{sequence_id}', '{genus_name}', '{genus_cn_name}', '{typ_name}', '{primer_name}', '{sequence_desc}', '{sequence}')"
        cursor.execute(sql)
        conn.commit()
