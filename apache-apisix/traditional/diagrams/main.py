from diagrams import Cluster, Diagram
from diagrams.onprem.compute import Server
from diagrams.onprem.client import Client

with Diagram("Apache APISIX detailed analysis", show=False):
    api_client = Client("API Client")

    with Cluster("Apache Gateway"):
        router = Server("Router")
        rate_limiter = Server("Rate Limiter")
        load_balancer = Server("Load Balancer")

    with Cluster("Services"):
        with Cluster("Quotes Service"):
            quotes = Server("quotes")
        with Cluster("Randomizer Service"):
            random_one = Server("random-one")
            random_two = Server("random-two")
        with Cluster("Nginx Services"):
            web_one = Server("Nginx-one")
            web_two = Server("Nginx-two")
        with Cluster("Jokes Service"):
            jokes = Server("jokes")

    api_client >> rate_limiter
    rate_limiter >> router
    router >> quotes
    router >> load_balancer
    router >> load_balancer
    load_balancer >> random_one
    load_balancer >> random_two
    load_balancer >> web_one
    load_balancer >> web_two
    router >> jokes