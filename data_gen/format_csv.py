import re
import csv
import os

# 定数として定義する
FILENAME = "donguri.csv"
DIR = "data"
PATTERN = "\d+"

# ファイルの存在確認と削除を関数として切り出す
def delete_file_if_exists(filename):
    if os.path.exists(filename):
        os.remove(filename)

# csvファイルの読み書きを関数として切り出す
def write_csv(filename, data):
    with open(filename, "w") as f:
        writer = csv.writer(f)
        writer.writerow(["episode", "script"])
        for episode, content in data:
            writer.writerow([episode, content])

def read_csv(filename):
    with open(filename, newline="") as f:
        reader = csv.reader(f)
        next(reader) # ヘッダー行をスキップする
        return list(reader)

# ファイルの存在確認と削除を実行する
delete_file_if_exists(FILENAME)

# データディレクトリからファイル名のリストを取得する
files = os.listdir(DIR)

# データを格納するリストを作成する
data = []

# 各ファイルに対して処理を行う
for file in files:
    # ファイル名からエピソード番号を取得する
    number = re.search(PATTERN, file)
    episode = int(number.group())
    # ファイルのパスを作成する
    path = os.path.join(DIR, file)
    # csvファイルを読み込む
    lines = read_csv(path)
    # 最後の行の内容を取得する
    content = lines[-1][0]
    # データに追加する
    data.append((episode, content))

# csvファイルに書き込む
write_csv(FILENAME, data)
