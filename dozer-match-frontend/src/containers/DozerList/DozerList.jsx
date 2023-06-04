import React from 'react';
import Dozer from '../../components/Dozer/Dozer';

const DozerList = ({cardData}) => {
  return (
    <div className="card-list">
      {cardData.map((card) => (
        <Dozer key={card.id} {...card} />
      ))}
    </div>
  );
};

export default DozerList;