import sys
import os
import glob
from PIL import Image

def strip_metadata(input_path, output_path, width, height):
    with Image.open(input_path) as img:
        clean = Image.new(img.mode, img.size)
        clean.putdata(list(img.get_flattened_data()))
        clean = clean.resize((width, height), Image.Resampling.LANCZOS)
        clean.save(output_path)
    print(f"CLEANED: {input_path} -> {output_path}")

if __name__ == "__main__":
    if len(sys.argv) < 2:
        myname = os.path.basename(__file__)
        print(f"画像ファイルからメタ情報を削除します")
        print(f"Usage: {myname} <file1> ...")
        exit(0)

    paths = []
    for arg in sys.argv[1:]:
        paths.extend(glob.glob(arg) or [arg])

    for path in paths:
        try:
            ifile = path
            ofile = "_" + path
            strip_metadata(ifile, ofile, 256, 256)
        except Exception as e:
            print(f"ERROR: {e}")
