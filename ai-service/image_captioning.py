import torch
from lavis.models import load_model_and_preprocess
from PIL import Image

# Device configuration
device = torch.device("cuda" if torch.cuda.is_available() else "cpu")

# Load BLIP captioning model and preprocessors
model, vis_processors, txt_processors = load_model_and_preprocess(
    name="blip_caption", 
    model_type="base_coco", 
    is_eval=True, 
    device=device
)

def generate_caption(image_path):
    """
    Generate a caption for the given image.

    Args:
        image_path (str): Path to the image file.

    Returns:
        str: Generated caption for the image.
    """
    # Load and preprocess the image
    raw_image = Image.open(image_path).convert("RGB")
    image = vis_processors["eval"](raw_image).unsqueeze(0).to(device)

    # Generate caption
    with torch.no_grad():
        caption = model.generate({"image": image})

    return caption[0]

# Example usage
# image_path = "park.jpg"
# caption = generate_caption(image_path)
# print("Generated Caption:", caption)
