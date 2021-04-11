import time
import random
from locust import HttpUser, task, between

class TrackerUser(HttpUser):
    wait_time = between(0.1, 0.2)

    @task
    def trail(self):
        self.client.post("/", json={"account_id":random.randint(1,10), "data":"short demo text"})
