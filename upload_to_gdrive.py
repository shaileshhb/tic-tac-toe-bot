import os
import json
from datetime import datetime
from google.oauth2 import service_account
from googleapiclient.discovery import build
from googleapiclient.http import MediaFileUpload

# Write the service account JSON from environment variable to a temp file
with open("sa.json", "w") as f:
    f.write(os.environ["GOOGLE_CREDENTIALS"])

# Load credentials from the temp file
creds = service_account.Credentials.from_service_account_file(
    "sa.json",
    scopes=["https://www.googleapis.com/auth/drive.file"],
)

# Build the Google Drive API service
drive_service = build("drive", "v3", credentials=creds)

# Rename the file with today's date
today = datetime.now().strftime("%Y%m%d")
file_name = f"myapp-{today}"

if os.path.isfile("myapp"):
    print("yes it exists")
else:
    print("no it does not exist")

# Google Drive folder ID (replace with your actual folder ID)
FOLDER_ID = os.environ["FOLDER_ID"]

# Upload the build file
file_metadata = {
    "name": file_name,
    "parents": [FOLDER_ID]
}
media = MediaFileUpload("myapp", resumable=True)

file = drive_service.files().create(
    body=file_metadata,
    media_body=media,
    fields="id"
).execute()

print(f"Uploaded File ID: {file.get('id')}")
