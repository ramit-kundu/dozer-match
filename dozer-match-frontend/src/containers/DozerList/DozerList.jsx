import React from 'react';
import Dozer from '../../components/Dozer/Dozer';


const cardData = [
  {
    id: 1,
    category: 'Excavator',
    engine: 'Diesel',
    hp: '200',
    weight: '5000kg',
    make: 'ABC Company',
    model: 'XYZ Model',
    imageUrl: 'https://s7d2.scene7.com/is/image/Caterpillar/CM20200429-439d6-30cb1',
  },
  {
    id: 2,
    category: 'Bulldozer',
    engine: 'Diesel',
    hp: '300',
    weight: '7000kg',
    make: 'DEF Company',
    model: 'PQR Model',
    imageUrl: 'https://s7d2.scene7.com/is/image/Caterpillar/CM20200429-439d6-30cb1',
  },
  // Add more card objects as needed
];

const DozerList = () => {
  return (
    <div className="card-list">
      {cardData.map((card) => (
        <Dozer key={card.id} {...card} />
      ))}
    </div>
  );
};

export default DozerList;