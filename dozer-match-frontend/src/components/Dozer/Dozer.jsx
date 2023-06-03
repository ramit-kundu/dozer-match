import React from 'react';
import { Card as AntdCard } from 'antd';
import './Dozer.css'

const { Meta } = AntdCard;

const Dozer = ({ category, engine, hp, weight, make, model, imageUrl }) => {
  return (
    <AntdCard
      hoverable
      cover={<img alt={make} src={imageUrl} />}
      className="card"
    >
      <Meta title={model} description={make} />
      <div>
        <p>Category: {category}</p>
        <p>Engine: {engine}</p>
        <p>HP: {hp}</p>
        <p>Operating Weight: {weight}</p>
      </div>
    </AntdCard>
  );
};

export default Dozer;