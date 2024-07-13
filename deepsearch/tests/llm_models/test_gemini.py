import os
import unittest

from src.llm_models.gemini import Gemini


class TestGemini(unittest.TestCase):
    # def test_gemini_key(self):
    #     """
    #     Test if the Gemini model has a key.
    #     """
    #     api = os.getenv("GOOGLE_API_KEY")
    #     print(api)
    #     self.assertTrue(api != "")
        
    def test_gemini_connection(self):
        """
        Test if the Gemini model can be connected to.
        """
        gemini = Gemini()
        response = gemini.query("What is the capital of France?", {})
        print(response)
        self.assertTrue(response["llm_response"] != "")
    
    # AIzaSyDIRWCtvPAH1XkUDIADSpkJUCnq0TIQ9U8