from Crypto.Util.Padding import pad, unpad
from Crypto.Cipher import AES
from Crypto.Random import get_random_bytes
from base64 import b64encode, b64decode

key = get_random_bytes(32)
iv = get_random_bytes(16)


def encrypt(plaintext, key):
    cipher = AES.new(key, AES.MODE_CBC, iv)
    return b64encode(cipher.encrypt(pad(plaintext.encode(), 16))).decode()


def decrypt(ciphertext, key):
    cipher = AES.new(key, AES.MODE_CBC, iv)
    return unpad(cipher.decrypt(b64decode(ciphertext.encode())), 16).decode()


def main():
    message = input("Ingrese el mensaje a cifrar: ")
    encrypted = encrypt(message, key)

    print("Mensaje encriptado: ", str(encrypted))
    message_decrypted = decrypt(encrypted, key)
    print(" ")
    print("Mensaje desencriptado: ", str(message_decrypted))


if __name__ == "__main__":
    main()
