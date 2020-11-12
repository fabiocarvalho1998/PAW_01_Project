<template>
 <div>
   <h1>Coords:</h1>
   <p>{{ coordinates.lat }} Latitude, {{ coordinates.lng }} Longitude</p>
   <gmap-map 
        :center="{lat:coordinates.lat, lng:coordinates.lng}"
        :zoom="7"
        style="width:640px; height:360px; margin: 32px auto;"
        ref="mapRef"
    >
        <gmap-info-window
            :options="{
                pixelOffset: {
                    width: 0,
                    height: -35
                }
            }"
            :position="{lat:coordinates.lat, lng:coordinates.lng}"
            :opened="infoWindowOpened"
            @closeclick="infoWindowOpened = false"
        >
            <div>
                <h2>Evento de Teste</h2>
                <h5>Participantes: TO-DO</h5>
            </div>
        </gmap-info-window>
        <gmap-marker
            :position="{lat:coordinates.lat, lng:coordinates.lng}"
            :clickable="true"
            :draggable="false"
            @click="infoWindowOpened = true"
        >
        </gmap-marker>
    </gmap-map>
 </div>
</template>

<script>

export default {
  data() {
    return {
      infoWindowOpened: false,
      coordinates: {
          lat: 0,
          lng: 0
      }
    };
  },
  created() {
      // get current coords from browser
      this.$getLocation({})
        .then(coordinates => {
            this.coordinates = coordinates
        }).catch(error => {
            alert(error)
        })
  },
}
</script>