import React, { useEffect, useState } from 'react';

function CryptoPrice() {
  const [price, setPrice] = useState(null);

  useEffect(() => {
    const ws = new WebSocket('ws://localhost:8080/ws'); // Kết nối WebSocket tới backend

    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      setPrice(data.price); // Cập nhật giá mới
    };

    return () => ws.close(); // Đóng WebSocket khi component unmount
  }, []);

  return <div>Giá hiện tại: {price ? price : 'Loading...'}</div>;
}

export default CryptoPrice;
