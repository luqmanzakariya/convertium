<template>
  <div class="room-card">
    <div class="room-image">
      <img :src="room.imageUrl" />
    </div>
    <div class="room-details">
      <div class="room-info">
        <div class="room-info-title">{{ room.name }}</div>
        <div class="room-info-subtitle">{{ room.subtitle }}</div>
        <div class="room-info-description">{{ room.description }}</div>
      </div>
      <div class="room-price">
        <div class="price">
          S${{ room.price }}<span class="price-caption">/night</span>
        </div>
        <div class="additional-info">Subject to GST and charges</div>
        <button @click="$emit('book-room', room.id)">Book Room</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "RoomCard",
  props: {
    room: {
      type: Object,
      required: true,
    },
  },
  emits: ["book-room"],
};
</script>

<style scoped>
.room-card {
  display: flex;
  border: 1px solid #ddd;
  margin-bottom: 20px;
  background-color: #f0f1eb;
}
.room-image img {
  width: 340px;
  height: 210px;
  padding: 12px 12px 12px 12px;
}

.room-details {
  flex-grow: 1;
  display: flex;
}
.room-details h2 {
  margin: 0;
}
.room-info {
  flex-grow: 1;
  background-color: #f0f1eb;
  padding: 12px;
}
.room-info-title {
  font-size: 24px;
  font-weight: 700;
  text-transform: uppercase;
  margin-bottom: 4px;
}
.room-info-subtitle {
  font-size: 16px;
  color: #666;
  font-weight: 500;
  margin-bottom: 8px;
}
.room-info-description {
  font-size: 12px;
  font-weight: 400;
  color: black;
}
.room-price {
  min-width: 150px;
  background-color: #e1e2d6;
  padding: 12px 24px;
}
.price {
  font-weight: bold;
  font-size: 24px;
  font-weight: 400;
}
.price-caption {
  font-size: 14px;
  font-weight: 600;
  color: #666;
}
.additional-info {
  font-size: 11px;
  color: #666;
  margin-top: 8px;
}
button {
  margin-top: 32px;
  border-radius: 4px;
  text-transform: uppercase;
  background-color: #000000;
  border: none;
  color: white;
  padding: 8px 16px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 600;
}
</style>
