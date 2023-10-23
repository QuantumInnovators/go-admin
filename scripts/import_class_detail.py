import pandas as pd
import pymysql

# 配置MySQL连接信息
host = "localhost"
port = 3306
user = "root"
password = "raspberrypi"
database = "aquatic"

files = ["底栖动物.xlsx", "浮游动物.xlsx", "浮游植物.xlsx", "水生植物.xlsx", "鱼类名录.xlsx"]
# 读取Excel文件
excel_file = "/Users/zhouyaoxu/Desktop/参考设计/名录7.24/"

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
    # 遍历数据并插入MySQL表中
    # 插入所有界
    for index, row in df.iterrows():
        # 判断是否已经插入过
        if row[2] in inserted:
            continue
        try:
            # 解析数据
            column1 = row[2]  # 拉丁学名
            column2 = row[1]  # 中文名
            sql = f"INSERT INTO `kingdom` (`name`, `desc`, `category_id`) VALUES ('{column1}', '{column2}', '{category_index}')"
            cursor.execute(sql)
            conn.commit()
            last_id = cursor.lastrowid
            id_dict[column1] = last_id
            # 将插入过的数据加入set
            inserted.add(row[2])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error occurred during commit:", error)

    # 插入所有门
    for index, row in df.iterrows():
        if row[4] in inserted:
            continue
        try:
            column1 = row[4]
            column2 = row[3]
            key = row[2]  # 所属界学名
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
            inserted.add(row[4])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error occurred during commit:", error)

    for index, row in df.iterrows():
        if row[6] in inserted:
            continue
        try:
            column1 = row[6]
            column2 = row[5]
            key = row[4]
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
            inserted.add(row[6])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error occurred during commit:", error)

    for index, row in df.iterrows():
        if row[8] in inserted:
            continue
        try:
            column1 = row[8]
            column2 = row[7]
            key = row[6]
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
            inserted.add(row[8])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error occurred during commit:", error)

    for index, row in df.iterrows():
        if row[10] in inserted:
            continue
        try:
            column1 = row[10]
            column2 = row[9]
            key = row[8]
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
            inserted.add(row[10])
        except pymysql.connector.Error as error:
            # 发生错误时进行处理
            print("Error occurred during commit:", error)

    for index, row in df.iterrows():
        if row[12] in inserted:
            continue
        try:
            column1 = row[12]
            column2 = row[11]
            key = row[10]
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
            inserted.add(row[12])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error occurred during commit:", error)

    for index, row in df.iterrows():
        if row[14] in inserted:
            continue
        try:
            column1 = row[14]
            column2 = row[13]
            key = row[12]
            fam_id = 0
            if key in id_dict:
                fam_id = id_dict[key]
            else:
                print("Not Found Parent Error: " + key)
            sql = f"INSERT INTO `species` (`name`, `desc`, `genus_id`, `category_id`) VALUES ('{column1}', '{column2}', '{fam_id}', '{category_index}')"
            cursor.execute(sql)
            conn.commit()
            inserted.add(row[14])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error occurred during commit:", error)
    category_index += 1
# 关闭连接
cursor.close()
conn.close()
