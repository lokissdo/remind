import redis


class RedisManager:
    def __init__(self, host: str):
        self.redis = redis.from_url(host)

    def set(self, key, value, expiry):
        self.redis.set(key, value, ex=expiry)

    def get(self, key):
        return self.redis.get(key)