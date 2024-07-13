from typing import Dict, List
from llm_models.base import BaseLLM
from llm_models.configs.gemini import GeminiConfig
from src.types.dtypes import MediaData, QueryResult
from src.types.media_type import MEDIA_TYPE

from langchain_google_genai import ChatGoogleGenerativeAI
from langchain.chains.llm import LLMChain


class Gemini(BaseLLM):
    def __init__(self, config: GeminiConfig = GeminiConfig()):
        self.config = config
        super().__init__()

    def query(
        self,
        query: str,
        contexts: Dict[MEDIA_TYPE, List[MediaData]],
    ) -> QueryResult:
        results = []
        for item in contexts.items():
            media_data = item[1]
            for each_response in media_data:
                results.append(each_response.get("document", ""))
        contexts = (" | ").join(results)
        llm_response = self._get_answer(contexts, query, self.config)
        query_result = {"llm_response": llm_response, "documents": contexts}
        return query_result

    def _get_llm_model(self):
        return ChatGoogleGenerativeAI(model="gemini-pro")

    def _get_answer(self, context: str, query: str, config: GeminiConfig) -> str:
        llm_model = self._get_llm_model()
        llm_chain = LLMChain(llm=llm_model, prompt=self.DEFAULT_PROMPT_TEMPLATE, verbose=True)
        response = llm_chain.run(context=context, query=query)
        return response