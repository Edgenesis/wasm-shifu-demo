# shifu && WasmEdge
# 
```mermaid
  flowchart TD
    
   iot[IOT #mock Device]
   shifu1[Shifu #to get Info]
   wasm[WasmEdge Go Server # wash Data]
   app[Application # collect Data]

   iot -->|telemetry| shifu1 -->|push| wasm -->|http get| app 
   
```

```
   shifu2[Shifu #to Collector]<!-- shifu2 <-->|curl get_info| app -->
```

## Iot Output
```
{
   "statusCode":"200",
   "message":"success",
   "entity":[
      {
         "datetime":"2022-08-18 19:43:34",
         "eUnit":"℃",
         "eValue":"27.4",
         "eKey":"e3",
         "eName":"大气温度",
         "eNum":"101"
      },
      {
         "datetime":"2022-08-18 19:43:34",
         "eUnit":"%RH",
         "eValue":"82.5",
         "eKey":"e4",
         "eName":"大气湿度",
         "eNum":"102"
      },
   ],
   "deviceId":18000856,
   "deviceName":"18000856"
}
```
## result
```
[
   {
      "code":"20990922009",
      "name":"大气温度",
      "val":"37",
      "unit":"℃",
      "exception":"温度过高"
   },
   {
      "code":"20990922009",
      "name":"大气湿度",
      "val":"35",
      "unit":"%RH",
      "exception":"湿度过高"
   }
]
```

# How to run ?
## run mock device
```bash
docker build . -t testmockdevice:v0.0.1
docker run  -p 8080:8080 -itd testmockdevice:v0.0.1 
```
## run shifu
```
kubectl apply -f shifuConfig/shifu_install.yml
kubectl apply -f shifuConfig/Shifu1
```

## run wasmEdge
rules path: wasmEdge/js-func/src/js/run.js
> docker docker build . -t testwasm:latest -f wasmEdge/dockerfile
> docker run -p 8080:8080 -itd testwasm