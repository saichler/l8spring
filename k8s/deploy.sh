kubectl apply -f ./vnet.yaml
sleep 2
kubectl apply -f ./spring.yaml
kubectl apply -f ./web.yaml
