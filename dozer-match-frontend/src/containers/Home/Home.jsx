import React from 'react';
import { Layout, Row, Col } from 'antd';


import DozerList from '../DozerList/DozerList';
import MyForm from '../../components/Form/Form';

const { Content } = Layout;

const getDozerList = ()=>{
  return <Row>
    <Col xs={24} md={12} lg={12}>
    <DozerList/>
          </Col>
          <Col xs={24} md={12} lg={12}>
          <DozerList/>
          </Col>
  </Row>
}

const Home = () => {
  return (
    <Layout>
      <Content style={{ padding: '50px' }}>
        <Row>
        <Col xs={24} md={12} lg={12}>
          <MyForm/>
          </Col>
          <Col xs={24} md={12} lg={12}>
          {getDozerList()}
          </Col>
        </Row>
      </Content>
    </Layout>
  );
};

export default Home;
