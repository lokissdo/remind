import sys
import os
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), 'pb')))
import grpc
from concurrent import futures
import pb.ai_service_pb2 as ai_service_pb2
import pb.ai_service_pb2_grpc as ai_service_pb2_grpc
from stt import summarize_audio
from image_captioning import generate_caption
from gemini import reformat_journal_entry, generate_todos
import google.protobuf.empty_pb2 as empty_pb2
import io
from google.protobuf.json_format import MessageToJson
import json
import base64
class AIServiceServicer(ai_service_pb2_grpc.AIServiceServicer):
    def GenerateCaption(self, request, context):
        print("Received GenerateCaption request")
        try:
            image_data =  base64.b64decode(request.text)
            output_file_path = "output_image.png"

            # Write the decoded data to the file
            with open(output_file_path, "wb") as image_file:
                image_file.write(image_data)
            converted_image_file = open(output_file_path, "rb")
            caption = generate_caption(converted_image_file)
            print("Generated caption successfully")
            return ai_service_pb2.TextResponse(text=caption)
        except Exception as e:
            print(f"Error generating caption: {e}")
            context.set_details(f"Error generating caption: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            return ai_service_pb2.TextResponse()

    def ReformatJournalEntry(self, request, context):
        print("Received ReformatJournalEntry request")
        try:
            reformatted_text = reformat_journal_entry(request.text)
            print("Reformatted journal entry successfully")
            return ai_service_pb2.TextResponse(text=reformatted_text)
        except Exception as e:
            print(f"Error reformatting journal entry: {e}")
            context.set_details(f"Error reformatting journal entry: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            return ai_service_pb2.TextResponse()

    def GenerateTodos(self, request, context):
        print("Received GenerateTodos request")
        try:
            todos = generate_todos(request.text)
            print("Generated todos successfully")
            return ai_service_pb2.TextResponse(text=todos)
        except Exception as e:
            print(f"Error generating todos: {e}")
            context.set_details(f"Error generating todos: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            return ai_service_pb2.TextResponse()

    def SummarizeAudio(self, request, context):
        print("Received SummarizeAudio request")
        try:
            audio_data =  base64.b64decode(request.text)
            output_file_path = "output_audio.m4a"

            # Write the decoded data to the file
            with open(output_file_path, "wb") as audio_file:
                audio_file.write(audio_data)
            summary = summarize_audio(output_file_path)
            print("Summarized audio successfully")
            return ai_service_pb2.TextResponse(text=summary)
        except Exception as e:
            print(f"Error summarizing audio: {e}")
            context.set_details(f"Error summarizing audio: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            return ai_service_pb2.TextResponse()

def serve():
    print("Starting AI server")
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    ai_service_pb2_grpc.add_AIServiceServicer_to_server(AIServiceServicer(), server)
    server.add_insecure_port('[::]:5006')
    server.start()
    print("AI server started on port 5006")
    server.wait_for_termination()

if __name__ == "__main__":
    serve()
