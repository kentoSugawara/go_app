import React, { useEffect, useState } from 'react';
import axios from 'axios';

function ListComponent() {
  const [items, setItems] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:8080/list')
      .then(response => {
        setItems(response.data);
      })
      .catch(error => {
        console.error('Error fetching data: ', error);
      })
  }, []);

  return (
    <div>
      {items.map(item => (
        <div key={item.URL}>
          <h2>{item.Title}</h2>
          <p>{item.Body}</p>
          <a href={item.URL}>Read more</a>
        </div>
      ))}
    </div>
  );
}

export default ListComponent;