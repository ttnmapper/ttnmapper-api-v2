import os
from celery import Celery

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'ttnm_backend.settings')

# TODO: Figure out how to set via config file
app = Celery('api', backend='amqp', broker='amqp://guest@localhost//')
app.config_from_object('django.conf:settings', namespace='CELERY')
app.autodiscover_tasks()
