import google.generativeai as genai
import json
from datetime import datetime
genai.configure(api_key="AIzaSyCSJaNW2VIaMiny5-q2W7MfN8wCNt6WrCY")

# Set up the model
generation_config = {
  "temperature": 0.9,
  "top_p": 1,
  "top_k": 1,
  "max_output_tokens": 2048,
}

safety_settings = [
  {
    "category": "HARM_CATEGORY_HARASSMENT",
    "threshold": "BLOCK_MEDIUM_AND_ABOVE"
  },
  {
    "category": "HARM_CATEGORY_HATE_SPEECH",
    "threshold": "BLOCK_MEDIUM_AND_ABOVE"
  },
  {
    "category": "HARM_CATEGORY_SEXUALLY_EXPLICIT",
    "threshold": "BLOCK_MEDIUM_AND_ABOVE"
  },
  {
    "category": "HARM_CATEGORY_DANGEROUS_CONTENT",
    "threshold": "BLOCK_MEDIUM_AND_ABOVE"
  }
]

gemini_model = genai.GenerativeModel(model_name="gemini-pro",
                              generation_config=generation_config,
                              safety_settings=safety_settings)


# example prompt
# prompt_parts = [
#     """
#     Give me code for adversarial attack in python
#     """
# ]

# response = gemini_model.generate_content(prompt_parts)
# print(response.text)
def generate_todos(journal_entry):
    """
    Generate a list of todos based on the given journal entry.

    Args:
        journal_entry (str): Journal entry text.

    Returns:
        list: List of todos generated from the journal entry.
    """
    # Generate todos
    prompt_parts = [
        journal_entry+
        """
        \nFrom this journal , generate todo IN THE FUTURE that MUST FOLLOW THIS FORMAT: 
        {
            "todos": [
                {
                    "title": "Example Todo",
                    "description": "This is an example todo",
                    "startTime": "2024-06-10T09:00:00Z",
                    "endTime": "2024-06-15T10:00:00Z",
                    "userID": "your_user_id",
                    "reminders": [
                        {
                            "start": "2024-06-10T08:30:00Z",
                            "message": "Reminder 1"
                        },
                        {
                            "start": "2024-06-15T08:45:00Z",
                            "message": "Reminder 2"
                        }
                    ]
                },
            ]
        }

        Your response MUST INCLUDE ONLY JSON FORMAT
        """
        + "Today is" +datetime.now().date().isoformat()
    ]

    response = gemini_model.generate_content(prompt_parts)
    
    # Check if reponse is in JSON format
    try:
        todos = json.loads(response.text)
        return response.text
    except json.JSONDecodeError:
        todos = {"error": "Failed to generate todos. Please try again."}
        return todos
   


# print(generate_todos(journal_entry))


def reformat_journal_entry(journal_entry):
    """
    Reformat the journal entry text for better readability.

    Args:
        journal_entry (str): Journal entry text.

    Returns:
        str: Reformatted journal entry text.
    """
    # Reformat journal entry
    prompt_parts = [
        f"""
        {journal_entry}
        Reformat the journal entry text for correct grammar.
        Your response MUST INCLUDE ONLY FORMATTED JOURNAL
        """
    ]

    response = gemini_model.generate_content(prompt_parts)
    return response.text

def text_journal(text) :
    prompt_parts = [
        f"""
        {text}
        Give me an intimate journal paragraph based following information.
        Your response MUST INCLUDE ONLY JOURNAL
        """
    ]
    print(prompt_parts)
    response = gemini_model.generate_content(prompt_parts)
    return response.text
# print(reformat_journal_entry(journal_entry))

if __name__ == "__main__":
    test_journal_entry = """
    Today unfolded with a blend of productivity and creativity, beginning promptly at 9:00 AM with a revitalizing morning exercise session. The weather was ideal, providing a serene backdrop for a brisk walk followed by a series of stretching exercises in the park. This invigorating start set a positive tone for the day. By 10:30 AM, I found myself back home, savoring a nutritious breakfast of oatmeal, fresh fruits, and green tea. This quiet time allowed me to meticulously plan the day's activities, ensuring a structured and focused approach.

    At 11:00 AM, I immersed myself in an SEO project for a new client. The task demanded my full attention as I engaged in thorough keyword research and competitor analysis, laying the groundwork for a robust content strategy. The morning's work flowed seamlessly into the early afternoon, culminating in a satisfying sense of progress.

    Lunch at 1:00 PM provided a welcome respite. I prepared a delicious salad with grilled chicken, which I enjoyed while indulging in a few chapters of my current book. This brief interlude offered a refreshing break, recharging my mental and physical energies.

    The afternoon brought a shift in focus to my YouTube channel. At 2:00 PM, I embarked on a creative journey, brainstorming and recording engaging stories about farm animals. This process was both enjoyable and fulfilling, allowing me to explore new ideas and connect with my audience in a meaningful way.

    Reflecting on the day's events, I am struck by the harmonious balance between productivity and creativity. Each activity, from the morning exercise to the afternoon content creation, contributed to a well-rounded and rewarding day. Looking ahead to tomorrow, I anticipate another busy yet exciting day. I plan to start the morning at 8:30 AM with a yoga session to enhance my flexibility and mindfulness. By 10:00 AM, I will continue working on the SEO project, focusing on optimizing the content strategy. Lunch at 12:30 PM will be followed by a visit to the local library at 1:30 PM to gather research materials for my upcoming YouTube script. The afternoon, starting at 3:00 PM, will be dedicated to drafting and refining the script, ensuring that it captures the essence of the stories I wish to share.

    """
    print(text_journal("Today was a good day. I went to the park and had a picnic with my friends. We played games and enjoyed the sunny weather. I felt happy and relaxed."))
    print(reformat_journal_entry(test_journal_entry))
    print(generate_todos(test_journal_entry))