from pydub import AudioSegment
import deepspeech
# import wave
import numpy as np
import os
from config import *

BEAM_WIDTH = 500
AudioSegment.converter  = FFMPEG_PATH


sttModel = deepspeech.Model(STT_MODEL_PATH)
sttModel.enableExternalScorer(LM_PATH)
# sttModel.setScorerBeamWidth(BEAM_WIDTH)


def convert_audio(file):
    # Load audio file
    audio = AudioSegment.from_mp3(file)

    # Convert stereo to mono
    audio = audio.set_channels(1)

    # Change sample rate to 16000Hz
    audio = audio.set_frame_rate(16000)

    return audio

def transcribe(file):
        audio = convert_audio(file)

        # Save the result
        buffer = audio.get_array_of_samples()
        # Not 64-bit RIFF file
        # Convert buffer to 16-bit int array for DeepSpeech
        data16 = np.frombuffer(buffer, dtype=np.int16)

        # Perform transcription
        text = sttModel.stt(data16)

        return text

mp3_file_path ="./test.mp3"
# load mp3 file using os and transcribe

if os.path.exists(mp3_file_path):
    print(transcribe(mp3_file_path))
else:
    print("Audio file not found.")

