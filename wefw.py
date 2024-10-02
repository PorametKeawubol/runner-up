import hashlib  # นำเข้าห้องสมุด hashlib

# แมปตัวอักษรกับอีโมจิ
emoji_to_char = {
    '😀': 'A', '😃': 'B', '😄': 'C', '😁': 'D', '😆': 'E', 
    '😅': 'F', '😂': 'G', '🤣': 'H', '😊': 'I', '😇': 'J', 
    '🙂': 'K', '🙃': 'L', '😉': 'M', '😌': 'N', '😍': 'O', 
    '😘': 'P', '😗': 'Q', '😙': 'R', '😚': 'S', '🤗': 'T', 
    '🤩': 'U', '🤔': 'V', '🤨': 'W', '😐': 'X', '😑': 'Y', '😶': 'Z'
}

def emoji_decode(encoded_string):
    decoded_string = ''
    for emoji in encoded_string:
        if emoji in emoji_to_char:
            decoded_string += emoji_to_char[emoji]
        else:
            decoded_string += emoji  # เก็บอักขระที่ไม่ใช่ emoji ไว้เหมือนเดิม
    return decoded_string

def generate_md5_hash(text):
    md5_hash = hashlib.md5(text.encode()).hexdigest()
    return md5_hash

# ข้อความเข้ารหัสด้วยอีโมจิ
encoded_message = '🤣😆🙃🙃😍 😆😉😍😇😊 😅😍😙 😅🤩😌😌😑 🤣😀🤣😀🤣😀'

# ถอดรหัสข้อความ
decoded_message = emoji_decode(encoded_message)
print("Decoded Message:", decoded_message)

# สร้าง MD5 hash จากข้อความที่ถอดรหัส
md5_result = generate_md5_hash(decoded_message)

# กำหนด flag
flag = f"THCTT24{{{md5_result}}}"
print(f"Flag found: {flag}")
