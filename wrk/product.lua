wrk.method = "GET"
wrk.headers["Content-Type"] = "application/json"
wrk.headers["Authorization"] = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50Ijoia3VydDY3ODMiLCJ1c2VySWQiOjIsImNvZGUiOiJtUDlCSiIsImV4cCI6MTc3OTQyOTgzOH0.rKRjl1ty0Jn5eNKWo3tbhUyjtAEpBUILGS6MlhVWaRQ"

-- wrk -t4 -c100 -d10s -s wrk/product.lua https://www.maplestoryexchange.com/product