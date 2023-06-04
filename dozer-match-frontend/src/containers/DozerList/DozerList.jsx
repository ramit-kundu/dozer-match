import React from 'react';
import Dozer from '../../components/Dozer/Dozer';
import {  Col } from 'antd';
const DozerList = ({cardData}) => {
  return (
    <>
      {cardData.map((card) => (
        <Col  lg={12}>
              <div className="card-list">
          <Dozer key={card.id} {...card} />
          </div>
        </Col>
        
      ))}
</>
  );
};

export default DozerList;