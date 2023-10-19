import pandas as pd
import pymysql

# 配置MySQL连接信息
host = "localhost"
port = 3306
user = "root"
password = "raspberrypi"
database = "aquatic"

files = ["底栖动物.xlsx", "浮游动物.xlsx", "水生植物.xlsx", "鱼类名录.xlsx"]
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

for file in files:
    print(file)
    df = pd.read_excel(excel_file + file)
    # 遍历数据并插入MySQL表中
    for index, row in df.iterrows():
        # 判断是否已经插入过
        if row["拉丁学名"] in inserted:
            continue
        try:
            # 解析数据
            column1 = row["拉丁学名"]
            column2 = row["界"]
            sql = f"INSERT INTO `kingdom` (`name`, `desc`) VALUES ('{column1}', '{column2}')"
            cursor.execute(sql)
            conn.commit()
            last_id = cursor.lastrowid
            id_dict[column1] = last_id
            # 将插入过的数据加入set
            inserted.add(row["拉丁学名"])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error occurred during commit:", error)

    for index, row in df.iterrows():
        if row["拉丁学名.1"] in inserted:
            continue
        try:
            column1 = row["拉丁学名.1"]
            column2 = row["门"]
            key = row["拉丁学名"]
            sql = f"INSERT INTO `phylum` (`name`, `desc`, `kingdom_id`) VALUES ('{column1}', '{column2}', '{id_dict[key]}')"
            cursor.execute(sql)
            conn.commit()
            last_id = cursor.lastrowid
            id_dict[column1] = last_id
            inserted.add(row["拉丁学名.1"])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error occurred during commit:", error)

    for index, row in df.iterrows():
        if row["拉丁学名.2"] in inserted:
            continue
        try:
            column1 = row["拉丁学名.2"]
            column2 = row["纲"]
            key = row["拉丁学名.1"]
            sql = f"INSERT INTO `class` (`name`, `desc`, `phylum_id`) VALUES ('{column1}', '{column2}', '{id_dict[key]}')"
            cursor.execute(sql)
            conn.commit()
            last_id = cursor.lastrowid
            id_dict[column1] = last_id
            inserted.add(row["拉丁学名.2"])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error occurred during commit:", error)

    for index, row in df.iterrows():
        if row["拉丁学名.3"] in inserted:
            continue
        try:
            column1 = row["拉丁学名.3"]
            column2 = row["目"]
            key = row["拉丁学名.2"]
            sql = f"INSERT INTO `order` (`name`, `desc`, `class_id`) VALUES ('{column1}', '{column2}', '{id_dict[key]}')"
            cursor.execute(sql)
            conn.commit()
            last_id = cursor.lastrowid
            id_dict[column1] = last_id
            inserted.add(row["拉丁学名.3"])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error occurred during commit:", error)

    for index, row in df.iterrows():
        if row["拉丁学名.4"] in inserted:
            continue
        try:
            column1 = row["拉丁学名.4"]
            column2 = row["科"]
            key = row["拉丁学名.3"]
            sql = f"INSERT INTO `family` (`name`, `desc`, `order_id`) VALUES ('{column1}', '{column2}', '{id_dict[key]}')"
            cursor.execute(sql)
            conn.commit()
            last_id = cursor.lastrowid
            id_dict[column1] = last_id
            inserted.add(row["拉丁学名.4"])
        except pymysql.connector.Error as error:
            # 发生错误时进行处理
            print("Error occurred during commit:", error)

    for index, row in df.iterrows():
        if row["拉丁学名.5"] in inserted:
            continue
        try:
            column1 = row["拉丁学名.5"]
            column2 = row["属"]
            key = row["拉丁学名.4"]
            sql = f"INSERT INTO `genus` (`name`, `desc`, `family_id`) VALUES ('{column1}', '{column2}', '{id_dict[key]}')"
            cursor.execute(sql)
            conn.commit()
            last_id = cursor.lastrowid
            id_dict[column1] = last_id
            inserted.add(row["拉丁学名.5"])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error occurred during commit:", error)

    for index, row in df.iterrows():
        if row["拉丁学名.6"] in inserted:
            continue
        try:
            column1 = row["拉丁学名.6"]
            column2 = row["种"]
            key = row["拉丁学名.5"]
            sql = f"INSERT INTO `species` (`name`, `desc`, `genus_id`) VALUES ('{column1}', '{column2}', '{id_dict[key]}')"
            cursor.execute(sql)
            conn.commit()
            inserted.add(row["拉丁学名.6"])
        except pymysql.Error as error:
            # 发生错误时进行处理
            print("Error occurred during commit:", error)
# 关闭连接
cursor.close()
conn.close()
