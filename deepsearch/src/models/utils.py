from pathlib import Path
from typing import Any, Dict, List
import tiktoken


def chunk_data(data: str, chunk_size: int = 50) -> List[Dict[str, Any]]:
    encoding = tiktoken.get_encoding("cl100k_base")
    tokens = encoding.encode(data)
    chunks = []
    for i in range(0, len(tokens), chunk_size):
        chunk_tokens = tokens[i : i + chunk_size]
        chunk_text = encoding.decode(chunk_tokens)
        chunks.append(chunk_text)
    return chunks


def read_file(file_path: str) -> str:
    path = Path(file_path)
    if path.is_file() and path.suffix == ".txt":
        return path.read_text(encoding="utf-8")

    err_msg = f"File not found or unsupported file type: {file_path}"
    raise ValueError(err_msg)
