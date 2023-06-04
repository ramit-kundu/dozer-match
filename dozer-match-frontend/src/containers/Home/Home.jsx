import React, { useEffect, useState } from 'react';
import { Layout, Row, Col } from 'antd';

import {LoadNewScrape,GetNewScrape,GetScrape,UpdateFilter} from './Parser'

import DozerList from '../DozerList/DozerList';
import MyForm from '../../components/Form/Form';

const { Content } = Layout;




const Home = () => {

  const [category, setCategory] = useState([]);
  const [minEngineHP, setMinEngineHP] = useState(0);
  const [maxEngineHP, setMaxEngineHP] = useState(99999999);
  const [minOperatingWT, setMinOperatingWT] = useState(0);
  const [maxOperatingWT, setMaxOperatingWT] = useState(9999999);
  const [selectedCategory, setSelectedCategory] = useState([0,0]);
  const [selectedOperatingWt, setSelectedOperatingWt] = useState([0,0]);
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

  useEffect(() => {
    GetScrape(setScrapeData);
    },[]); 

  useEffect(() => {

    if(scrapeData !== undefined && scrapeData !== null && scrapeData.length !== 0 ){
      LoadNewScrape(scrapeData,selectedCategory,selectedOperatingWt,selectedEngineHP,setMinOperatingWT,setMaxOperatingWT,setCategory,setMinEngineHP,setMaxEngineHP, setSelectedCategory,setSelectedOperatingWt,setSelectedEngineHP, setDisplayData)
    }
      // eslint-disable-next-line react-hooks/exhaustive-deps
    },[scrapeData]); 

    useEffect(() => {
      if(scrapeData !== undefined && scrapeData !== null  && 
        selectedCategory !== undefined && selectedCategory !== null && selectedCategory.length !== 0  && 
        selectedOperatingWt !== undefined && selectedOperatingWt !== null && selectedOperatingWt.length === 2 && 
        selectedEngineHP !== undefined && selectedEngineHP !== null && selectedEngineHP.length === 2
        )
     setDisplayData([...UpdateFilter(scrapeData,selectedCategory,selectedOperatingWt,selectedEngineHP,setDisplayData)])
      // eslint-disable-next-line react-hooks/exhaustive-deps
      },[selectedCategory,selectedOperatingWt,selectedEngineHP]); 

  const handleCheckboxChange = (checkedValues) => {
     setSelectedCategory([...checkedValues]);
  };

  const handleEngineHPChange = (hp) => {
    setSelectedEngineHP([...hp])  
  };

  const handleOperatingWtChange = (wt) => {
    setSelectedOperatingWt([...wt])
  };

  const onRefresh = () => {
    GetNewScrape(setScrapeData)
  };

  const getDozerList = ()=>{
    return <Row>
      <DozerList xs={24} md={12} lg={12} cardData={displayData}/>
    </Row>
  }

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
