from locust import HttpUser, between, task


class WebsiteUser(HttpUser):
    wait_time = between(5, 15)
    
    # def on_start(self):
    #     self.client.post("/login", {
    #         "username": "test_user",
    #         "password": ""
    #     })
    
    @task
    def index(self):
        # self.client.get("/timer")
        self.client.get("/package")
        
    # @task
    # def about(self):
    #     self.client.get("/about/")