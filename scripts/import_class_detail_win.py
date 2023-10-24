import pandas as pd
import pymysql

# 配置MySQL连接信息
host = "192.168.2.160"
port = 3306
user = "root"
password = "123456"
database = "aquatic"

files = ["底栖动物.xlsx", "浮游动物.xlsx", "浮游植物.xlsx", "水生植物.xlsx", "鱼类名录.xlsx"]
# 读取Excel文件
excel_file = r"D:\资料库\工作\项目\数据库系统\资料\名录7.24\\"



# 连接MySQL数据库
conn = pymysql.connect(
    host=host, port=port, user=user, password=password, database=database
)
cursor = conn.cursor()
# 构建一个set
inserted = set()
# 构建一个dict 用于存储id
id_dict = {}
category_index = 1

for file in files:
    print("current:" + str(category_index))
    df = pd.read_excel(excel_file + file)
    # 数据清洗：填充空值为指定值（例如空字符串）并去除前后空格
    df.fillna('', inplace=True)  # 填充空值为''
    df = df.applymap(lambda x: x.strip() if isinstance(x, str) else x)  # 去除字符串前后空格
    # 遍历数据并插入MySQL表中
    # 插入所有界
    for index, row in df.iterrows():
        # 判断是否已经插入过
        if row['拉丁学名'] in inserted:
            continue
        try:
            # 解析数据
            column1 = row['拉丁学名']  # 拉丁学名
            column2 = row['界']  # 中文名
            sql = f"INSERT INTO `kingdom` (`name`, `desc`, `category_id`) VALUES ('{column1}', '{column2}', '{category_index}')"
            cursor.execute(sql)
            conn.commit()
            last_id = cursor.lastrowid
            id_dict[column1] = last_id
            # 将插入过的数据加入set
            inserted.add(row['拉丁学名'])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error [界:kingdom] occurred during commit:", error)

    # 插入所有门
    for index, row in df.iterrows():
        if row['拉丁学名.1'] in inserted:
            continue
        try:
            column1 = row['拉丁学名.1']
            column2 = row['门']
            key = row['拉丁学名']  # 所属界学名
            fam_id = 0
            if key in id_dict:
                fam_id = id_dict[key]
            else:
                print("Not Found Parent Error: " + key)
            sql = f"INSERT INTO `phylum` (`name`, `desc`, `kingdom_id`, `category_id`) VALUES ('{column1}', '{column2}', '{fam_id}', '{category_index}')"
            cursor.execute(sql)
            conn.commit()
            last_id = cursor.lastrowid
            id_dict[column1] = last_id
            inserted.add(row['拉丁学名.1'])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error [门:phylum] occurred during commit:", error)

    for index, row in df.iterrows():
        if row['拉丁学名.2'] in inserted:
            continue
        try:
            column1 = row['拉丁学名.2']
            column2 = row['纲']
            key = row['拉丁学名.1']
            fam_id = 0
            if key in id_dict:
                fam_id = id_dict[key]
            else:
                print("Not Found Parent Error: " + key)
            sql = f"INSERT INTO `class` (`name`, `desc`, `phylum_id`, `category_id`) VALUES ('{column1}', '{column2}', '{fam_id}', '{category_index}')"
            cursor.execute(sql)
            conn.commit()
            last_id = cursor.lastrowid
            id_dict[column1] = last_id
            inserted.add(row['拉丁学名.2'])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error [纲:class]occurred during commit:", error)

    for index, row in df.iterrows():
        if row['拉丁学名.3'] in inserted:
            continue
        try:
            column1 = row['拉丁学名.3']
            column2 = row['目']
            key = row['拉丁学名.2']
            fam_id = 0
            if key in id_dict:
                fam_id = id_dict[key]
            else:
                print("Not Found Parent Error: " + key)
            sql = f"INSERT INTO `order` (`name`, `desc`, `class_id`, `category_id`) VALUES ('{column1}', '{column2}', '{fam_id}', '{category_index}')"
            cursor.execute(sql)
            conn.commit()
            last_id = cursor.lastrowid
            id_dict[column1] = last_id
            inserted.add(row['拉丁学名.3'])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error [目: order] occurred during commit:", error)

    for index, row in df.iterrows():
        if row['拉丁学名.4'] in inserted:
            continue
        try:
            column1 = row['拉丁学名.4']
            column2 = row['科']
            key = row['拉丁学名.3']
            fam_id = 0
            if key in id_dict:
                fam_id = id_dict[key]
            else:
                print("Not Found Parent Error: " + key)
            sql = f"INSERT INTO `family` (`name`, `desc`, `order_id`, `category_id`) VALUES ('{column1}', '{column2}', '{fam_id}', '{category_index}')"
            cursor.execute(sql)
            conn.commit()
            last_id = cursor.lastrowid
            id_dict[column1] = last_id
            inserted.add(row['拉丁学名.4'])
        except pymysql.connect.Error as error:
            # 发生错误时进行处理
            print("Error [科:family] occurred during commit:", error)

    for index, row in df.iterrows():
        if row['拉丁学名.5'] in inserted:
            continue
        try:
            column1 = row['拉丁学名.5']
            column2 = row['属']
            key = row['拉丁学名.4']
            fam_id = 0
            if key in id_dict:
                fam_id = id_dict[key]
            else:
                print("Not Found Parent Error: " + key)
            sql = f"INSERT INTO `genus` (`name`, `desc`, `family_id`, `category_id`) VALUES ('{column1}', '{column2}', '{fam_id}', '{category_index}')"
            cursor.execute(sql)
            conn.commit()
            last_id = cursor.lastrowid
            id_dict[column1] = last_id
            inserted.add(row['拉丁学名.5'])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error [属:genus] occurred during commit:", error)

    for index, row in df.iterrows():
        if row['拉丁学名.6'] in inserted:
            continue
        try:
            column1 = row['拉丁学名.6']
            column2 = row['种']
            key = row['拉丁学名.5']
            fam_id = 0
            if key in id_dict:
                fam_id = id_dict[key]
            else:
                print("Not Found Parent Error: " + key)
            sql = f"INSERT INTO `species` (`name`, `desc`, `genus_id`, `category_id`) VALUES ('{column1}', '{column2}', '{fam_id}', '{category_index}')"
            cursor.execute(sql)
            conn.commit()
            inserted.add(row['拉丁学名.6'])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error [种:species] occurred during commit:", error)
    category_index += 1
# 关闭连接
cursor.close()
conn.close()
