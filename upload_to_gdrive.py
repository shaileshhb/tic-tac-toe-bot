import os, sys, json, datetime
from google.oauth2.credentials import Credentials
from googleapiclient.discovery import build
from googleapiclient.http import MediaFileUpload

if len(sys.argv) < 2:
    print("Usage: python upload_to_gdrive.py <file>")
    exit(1)

file_path = sys.argv[1]

# Load OAuth creds
creds_dict = json.loads(os.environ["GOOGLE_CREDENTIALS"])
creds = Credentials.from_authorized_user_info(creds_dict, scopes=["https://www.googleapis.com/auth/drive.file"])

drive_service = build("drive", "v3", credentials=creds)

# Folder ID from your "build-files"
FOLDER_ID = os.environ["FOLDER_ID"]

today = datetime.datetime.now().strftime("%Y-%m-%d")
file_metadata = {
    "name": f"{today}-{os.path.basename(file_path)}",
    "parents": [FOLDER_ID]
}
media = MediaFileUpload(file_path, resumable=True)

file = drive_service.files().create(body=file_metadata, media_body=media, fields="id").execute()
print(f"Uploaded {file_path} to Google Drive with ID: {file.get('id')}")
