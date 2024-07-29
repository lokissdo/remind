import base64

# Base64 string (example, replace with your actual base64 string)
base64_string = "your_base64_encoded_string_here"

# Decode the base64 string
image_data = base64.b64decode(base64_string)

# Specify the output file path
output_file_path = "output_image.png"

# Write the decoded data to the file
with open(output_file_path, "wb") as image_file:
    image_file.write(image_data)

print(f"Image saved to {output_file_path}")
