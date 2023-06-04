import React, { useEffect, useState } from 'react';
import { Layout, Row, Col } from 'antd';
import axios from 'axios'


import DozerList from '../DozerList/DozerList';
import MyForm from '../../components/Form/Form';

const { Content } = Layout;

const onFinish = (values) => {
  console.log('Form values:', values);
};




const Home = () => {

  const [category, setCategory] = useState([]);
  const [minEngineHP, setMinEngineHP] = useState(0);
  const [maxEngineHP, setMaxEngineHP] = useState(100);
  const [minOperatingWT, setMinOperatingWT] = useState(0);
  const [maxOperatingWT, setMaxOperatingWT] = useState(100);
  const [selectedCategory, setSelectedCategory] = useState([]);
  const [selectedOperatingWt, setSelectedOperatingWt] = useState([0,100]);
  const [selectedEngineHP, setSelectedEngineHP] = useState([0,100]);
  const [scrapeData,setScrapeData] = useState([])
  const [displayData,setDisplayData] = useState([])

  const getFormObject = ()=> {
    return {
      category : category,
      minEngineHP:minEngineHP,
      maxEngineHP:maxEngineHP,
      minOperatingWT:minOperatingWT,
      maxOperatingWT:maxOperatingWT,
      selectedCategory:selectedCategory,
      selectedOperatingWt:selectedOperatingWt,
      selectedEngineHP:selectedEngineHP,
      handleCheckboxChange: handleCheckboxChange,
      handleEngineHPChange:handleEngineHPChange,
      handleOperatingWtChange:handleOperatingWtChange,
      onRefresh : onRefresh
    }
  }

  // category,
  // minEngineHP,
  // maxEngineHP,
  // minOperatingWT,
  // maxOperatingWT,
  // selectedCategory ,
  // selectedOperatingWt ,
  // selectedEngineHP,
  // handleCheckboxChange ,
  // handleEngineHPChange,
  // handleOperatingWtChange ,
  // onRefresh

  const handleCheckboxChange = (checkedValues) => {
    setSelectedCategory(checkedValues);
  };

  const handleEngineHPChange = (minHP, maxHP) => {
    setMinEngineHP(minHP);
    setMaxEngineHP(maxHP);
  };

  const handleOperatingWtChange = (minWT, maxWT) => {
    setMinOperatingWT(minWT);
    setMaxOperatingWT(maxWT);
  };

  const onRefresh = () => {
    console.log("REFRESH")
  };




  const fetchScrapeData = async () => {
    try {
      const response = await axios.get('localhost:8002/scrape');
      if(response.status ==200){
        setScrapeData(response.data)
      }else {
        await fetchNewScrapeData()
      }
    } catch (error) {
      console.error(error);
    }
  };
  
  const fetchNewScrapeData = async () => {
    try {
      const response = await axios.post('localhost:8002/scrape');
      if(response.status ==200){
        setScrapeData(response.data)
      }
    } catch (error) {
      console.error(error);
    }
  };


  // useEffect(() => {
  //   fetchScrapeData();
  //   }, []); 

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

  var d = {onSubmit: onFinish}
  return (
    <Layout>
      <Content style={{ padding: '50px' }}>
        <Row>
        <Col xs={24} md={12} lg={12}>
          <MyForm key={0}{...getFormObject()}/>
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
