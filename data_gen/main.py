# colabで実行する想定のPythonスクリプト

# !pip3 install -U openai-whisper
# !sudo apt update && sudo apt install ffmpeg

from google.colab import drive
drive.mount('/content/drive')

import whisper
import csv
import os

model = whisper.load_model("medium")
dir = '/content/drive/MyDrive/donguri001-300'
files = os.listdir(dir)
scripts = list(filter(lambda x: x.endswith('mp3') or x.endswith('m4a'), files))
scripts.sort()

for script in scripts:
  csv_name = script.replace(".mp3", ".csv").replace(".m4a", ".csv").replace(" ", "_")
  path = dir + "/" + csv_name
  is_file = os.path.isfile(path)
  if is_file:
    print(path + " exist\n")
  else:
    print("start generation: " + path)
    result = model.transcribe(dir+ "/" + script, language="ja")
    with open(path, 'w') as f:
        writer = csv.writer(f)
        writer.writerow(['script'])
        writer.writerow([result['text']])
