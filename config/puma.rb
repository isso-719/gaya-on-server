threads 5, 5
workers 2
environment ENV['RACK_ENV'] || "production"
preload_app!
plugin :tmp_restart