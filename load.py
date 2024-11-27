import requests
import random
import time

# Configuration
BASE_URL = "http://localhost:8080"  # Your Go application's base URL
INTERVAL = 1  # Time in seconds between each batch of requests

# List of API endpoints to simulate load for
ENDPOINTS = [
    "/user/signup",  # Simulates new user creation
    "/user/login",   # Simulates login attempts
    "/api/data",     # Simulates general API requests
    "/api/error",    # Simulates API failures (if applicable)
]

def hit_endpoint(endpoint, simulate_failure=False):
    """Sends requests to the given endpoint."""
    url = f"{BASE_URL}{endpoint}"
    try:
        if "signup" in endpoint:
            payload = {"username": f"user{random.randint(1, 1000)}", "password": "pass"}
            response = requests.post(url, json=payload)
        elif "login" in endpoint:
            payload = {"username": "user", "password": "pass"}
            response = requests.post(url, json=payload)
        else:
            response = requests.get(url)

        if simulate_failure or response.status_code >= 400:
            print(f"Error response from {url} - Status: {response.status_code}")
        else:
            print(f"Successful response from {url} - Status: {response.status_code}")
        
    except requests.RequestException as e:
        print(f"Request to {url} failed: {e}")

def simulate_load():
    """Simulates user traffic to trigger Prometheus metrics."""
    for _ in range(random.randint(5, 10)):  # Random batch size
        endpoint = random.choice(ENDPOINTS)
        simulate_failure = random.choice([False, True]) if "api/error" in endpoint else False
        hit_endpoint(endpoint, simulate_failure)
        time.sleep(random.uniform(0.1, 0.5))  # Short random delay

if __name__ == "__main__":
    while True:
        print("Applying synthetic load...")
        simulate_load()
        print(f"Sleeping for {INTERVAL} seconds...")
        time.sleep(INTERVAL)
