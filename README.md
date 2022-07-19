# Aachen Mobility Exporter
Prometheus exporter for Aachen's mobility dashboard.

https://verkehr.aachen.de

## Usage
```bash
docker run -e PORT=9090 -p 9090:9090 ghcr.io/menzerath/aachen-mobility-exporter:latest
```

```bash
$ curl http://localhost:9090/metrics

# HELP aachen_mobility_parking_carpark_free How many car-park parking spots are free.
# TYPE aachen_mobility_parking_carpark_free gauge
aachen_mobility_parking_carpark_free{description="APAG Parkhaus",id="865",latitude="6.045195",longitude="50.774804",name="Uniklinik",trend="up"} 1744
...

# HELP aachen_mobility_parking_carpark_load Percentage of how many car-park parking spots are taken.
# TYPE aachen_mobility_parking_carpark_load gauge
aachen_mobility_parking_carpark_load{description="APAG Parkhaus",id="865",latitude="6.045195",longitude="50.774804",name="Uniklinik",trend="up"} 14.55
...

# HELP aachen_mobility_parking_roadside_free How many roadside parking spots are free.
# TYPE aachen_mobility_parking_roadside_free gauge
aachen_mobility_parking_roadside_free{latitude="50.770016",location_id="233",longitude="6.099145",name="Friedrichstraße 111-117",parent_locations="1",sub_locations="0",type="PARKING_AREA"} 0
aachen_mobility_parking_roadside_free{latitude="50.772200",location_id="364",longitude="6.098808",name="Friedrichstraße Süd",parent_locations="1",sub_locations="4",type="GROUP"} 11
aachen_mobility_parking_roadside_free{latitude="50.772953",location_id="366",longitude="6.098711",name="Friedrichstraße",parent_locations="0",sub_locations="2",type="GROUP"} 13
...

# HELP aachen_mobility_parking_roadside_occupied How many roadside parking spots are occupied.
# TYPE aachen_mobility_parking_roadside_occupied gauge
aachen_mobility_parking_roadside_occupied{latitude="50.770016",location_id="233",longitude="6.099145",name="Friedrichstraße 111-117",parent_locations="1",sub_locations="0",type="PARKING_AREA"} 1
aachen_mobility_parking_roadside_occupied{latitude="50.772200",location_id="364",longitude="6.098808",name="Friedrichstraße Süd",parent_locations="1",sub_locations="4",type="GROUP"} 97.31
aachen_mobility_parking_roadside_occupied{latitude="50.772953",location_id="366",longitude="6.098711",name="Friedrichstraße",parent_locations="0",sub_locations="2",type="GROUP"} 149.57
...

# HELP aachen_mobility_parking_roadside_total How many roadside parking spots are available in total.
# TYPE aachen_mobility_parking_roadside_total gauge
aachen_mobility_parking_roadside_total{latitude="50.770016",location_id="233",longitude="6.099145",name="Friedrichstraße 111-117",parent_locations="1",sub_locations="0",type="PARKING_AREA"} 1
aachen_mobility_parking_roadside_total{latitude="50.772200",location_id="364",longitude="6.098808",name="Friedrichstraße Süd",parent_locations="1",sub_locations="4",type="GROUP"} 108.31
aachen_mobility_parking_roadside_total{latitude="50.772953",location_id="366",longitude="6.098711",name="Friedrichstraße",parent_locations="0",sub_locations="2",type="GROUP"} 162.57
...
```
