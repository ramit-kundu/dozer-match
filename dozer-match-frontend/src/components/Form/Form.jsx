import React from 'react';
import { Form, Input, Button, Checkbox, Slider } from 'antd';
import './Form.css'

const MyForm = () => {
  const onFinish = (values) => {
    console.log('Form values:', values);
  };

  return (
    <div className="form-container">
      <Form onFinish={onFinish} className="form">
        <Form.Item
          label="Category"
          name="category"
          rules={[{ required: true, message: 'Please select at least one category' }]}
        >
          <Checkbox.Group>
            <Checkbox value="smallDozer">Small Dozer</Checkbox>
            <Checkbox value="mediumDozer">Medium Dozer</Checkbox>
            <Checkbox value="largeDozer">Large Dozer</Checkbox>
            <Checkbox value="wheelDozer">Wheel Dozer</Checkbox>
          </Checkbox.Group>
        </Form.Item>
        <Form.Item
          label="Engine HP"
          name="engineHP"
          rules={[{ required: true, message: 'Please select the engine HP' }]}
        >
          <Slider range min={0} max={1000} />
        </Form.Item>
        <Form.Item
          label="Operating Weight"
          name="operatingWeight"
          rules={[{ required: true, message: 'Please select the operating weight' }]}
        >
          <Slider range min={0} max={50000} />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit">
            Submit
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};

export default MyForm;
