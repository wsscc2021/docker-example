#Application modules
from flask import Flask
from healthcheck import HealthCheck

# Create Flask App
app = Flask(__name__)

# HelloWorld API
@app.route("/")
def index():
    return "Hello World!", 200

# Healthcheck API
health = HealthCheck()
app.add_url_rule("/healthcheck", "healthcheck", view_func=lambda: health.run())

if __name__ == "__main__":
    RUN_OPTIONS = {"host": "0.0.0.0", "port": 5000, "threaded": True}
    app.run(**RUN_OPTIONS)