import base64
import os

with open("~/.ssh/id_rsa.files.freelancer.vanilla", "rb") as file:
    data = file.read()

env_var = str(base64.b64encode(data),encoding='utf8')
os.environ["ID_RSA_FILES_FREELANCER_VANILLA"] =env_var
print(env_var)
print("\n\n")

with open("~/.ssh/id_rsa.files.freelancer.vanilla.out", "wb") as file:
    file.write(base64.b64decode(bytes(os.environ["ID_RSA_FILES_FREELANCER_VANILLA"], encoding='utf8')))
