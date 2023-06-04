import React from 'react';
import { Form,Row,Col, Button, Checkbox, Slider } from 'antd';
import './Form.css'

const MyForm = ({
  category,
  minEngineHP,
  maxEngineHP,
  minOperatingWT,
  maxOperatingWT,
  selectedCategory ,
  selectedOperatingWt ,
  selectedEngineHP,
  handleCheckboxChange ,
  handleEngineHPChange,
  handleOperatingWtChange ,
  onRefresh}) => {
  if(!Array.isArray(category) || category === undefined) {
    category = []
  }

  return (
    <div className="form-container">
      <Form className="form">
        <Form.Item
          label="Category"
          name="category"
        >
          <Checkbox.Group onChange={handleCheckboxChange} value={selectedCategory}>
            {category.map((c)=>{
              return <Checkbox value={c}>{c}</Checkbox>
            })}
          </Checkbox.Group>
        </Form.Item>
        <Form.Item
          label="Engine HP"
          name="engineHP"
        >
          <Slider range min={minEngineHP} max={maxEngineHP} onChange={handleEngineHPChange} value={selectedEngineHP}/>
        </Form.Item>
        <Form.Item
          label="Operating Weight"
          name="operatingWeight"
        >
          <Slider range min={minOperatingWT} max={maxOperatingWT} onChange={handleOperatingWtChange} value={selectedOperatingWt} />
        </Form.Item>
        <Form.Item>
        <Form.Item>
        <Row></Row>
        <Row>
          <br/>
        </Row>
          <Row>
          <Col lg={12} md={12} sm={0} >
            </Col>
            <Col lg={12} md={12} >
            <Button type="primary" onClick={onRefresh} className='formbutton'>
            REFRESH
          </Button>
            </Col>
          </Row>
        </Form.Item>
          
        </Form.Item>
        
      </Form>
    </div>
  );
};

export default MyForm;
