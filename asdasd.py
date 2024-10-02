import hashlib

# แมปอีโมจิกับตัวอักษร
emoji_to_char = {
    '🥷😀': 'A', '🥷😃': 'B', '🥷😄': 'C', '🥷😁': 'D', '🥷😆': 'E',
    '🥷😅': 'F', '🥷😂': 'G', '🥷🤣': 'H', '🥷😊': 'I', '🥷😇': 'J',
    '🥷🙂': 'K', '🥷🙃': 'L', '🥷😉': 'M', '🥷😌': 'N', '🥷😍': 'O',
    '🥷😘': 'P', '🥷😗': 'Q', '🥷😙': 'R', '🥷😚': 'S', '🥷🤗': 'T',
    '🥷🤩': 'U', '🥷🤔': 'V', '🥷🤨': 'W', '🥷😐': 'X', '🥷😑': 'Y',
    '🥷😶': 'Z', '🥷🤖': '1', '🥷👽': '2', '🥷👾': '3', '🥷🤡': '4',
    '🥷🦸': '5', '🥷🦹': '6', '🥷👻': '7', '🥷💀': '8', '🥷👺': '9',
    '🥷👹': '0', '🥷👿': '!', '🥷🤠': '@', '🥷🥳': '#', '🥷🥺': '$',
    '🥷🤯': '%', '🥷😰': '^', '🥷😱': '&', '🥷😵': '*', '🥷😡': '(',
    '🥷🤬': ')', '🥷😈': '-', '🥷👋': '=', '🥷👏': '+', '🥷🙌': '[',
    '🥷👐': ']', '🥷🤲': '{', '🥷🙏': '}', '🥷💪': ':', '🥷🦵': ';',
    '🥷🦶': '"', '🥷👀': "'", '🥷👁': '<', '🥷👃': '>', '🥷👂': '?',
    '🥷👄': '/', '🥷🦷': '\\', '🥷🦴': '|', '🥷👅': ' ', '🥷👄': '.',
    '🥷👓': ',', '🥷🕶': '`', '🥷🥽': '~', '🥷🎩': 'a', '🥷🧢': 'b',
    '🥷🪖': 'c', '🥷👒': 'd', '🥷🎓': 'e', '🥷⛑': 'f', '🥷👑': 'g',
    '🥷👔': 'h', '🥷👕': 'i', '🥷👖': 'j', '🥷👗': 'k', '🥷👘': 'l',
    '🥷👚': 'm', '🥷👛': 'n', '🥷👜': 'o', '🥷👝': 'p', '🥷🥷': 'q',
    '🥷🎒': 'r', '🥷🩲': 's', '🥷🩳': 't', '🥷👠': 'u', '🥷👡': 'v',
    '🥷👢': 'w', '🥷👞': 'x', '🥷🥿': 'y', '🥷👟': 'z'
}

def emoji_decode(encoded_string):
    decoded_string = ''
    emoji_list = encoded_string.split(' ')
    
    for emoji in emoji_list:
        if emoji in emoji_to_char:
            decoded_string += emoji_to_char[emoji]
        else:
            decoded_string += '?'  # เก็บอักขระที่ไม่รู้จัก
    
    return decoded_string

def generate_md5_hash(text):
    md5_hash = hashlib.md5(text.encode()).hexdigest()
    return md5_hash

# ข้อความที่เข้ารหัสด้วยอีโมจิ
encoded_message = "🥷😆 🥷😁 🥷😅 🥷😄"  # เปลี่ยนเป็นข้อความที่คุณต้องการ

# ถอดรหัสข้อความ
decoded_message = emoji_decode(encoded_message)
print("Decoded Message:", decoded_message)

# สร้าง MD5 hash จากข้อความที่ถอดรหัส
md5_result = generate_md5_hash(decoded_message)

# สร้าง flag
flag = f"THCTT24{{{md5_result}}}"
print(f"Flag found: {flag}")
