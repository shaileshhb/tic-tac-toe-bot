import os, json
from datetime import datetime
from google.oauth2 import service_account
from googleapiclient.discovery import build
from googleapiclient.http import MediaFileUpload

# Load service account credentials from env variable
creds_info = json.loads(os.environ["GOOGLE_CREDENTIALS"])
creds = service_account.Credentials.from_service_account_info(
    creds_info,
    scopes=["https://www.googleapis.com/auth/drive.file"],
)

# Build the Drive API service
drive_service = build("drive", "v3", credentials=creds)

# Get today's date and set file name
today = datetime.now().strftime("%Y%m%d")
file_name = f"myapp-{today}"

# Upload the build file
file_metadata = {"name": file_name}
media = MediaFileUpload("myapp", resumable=True)

file = drive_service.files().create(
    body={**file_metadata, "parents": os.environ["FOLDER_ID"]},  # Replace with actual folder ID
    media_body=media,
    fields="id"
).execute()

print("Uploaded File ID:", file.get("id"))
