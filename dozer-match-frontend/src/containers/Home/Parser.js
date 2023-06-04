import axios from 'axios'

 function Parse(scrapeData) {
   return scrapeData.map((data,i)=> {
      return {
        id : i,
        category : data["Category"],
        hp : data["EngineHP"],
        weight : data["OperatingWeight"],
        make : data["Make"],
        model : data["Model"],
        imageUrl : data["Picture"]?data["Picture"]:"https://static.vecteezy.com/system/resources/previews/005/337/799/original/icon-image-not-found-free-vector.jpg",
      }
    })
}

const GetScrape = async (setScrapeData,setAlert) => {
    try {
      setAlert({ message: 'Fetching Old Scrape', type: 'info' });
      const response = await axios.get('http://localhost:8002/scrape');
      if(response.status === 200 && response.data.length > 0 ){
        setAlert({ message: 'Fetched Old Scrape ', type: 'success' });
        setScrapeData(Parse(response.data))
        return;
      }else {
        setAlert({ message: 'No existing scrape found trying to fetch new scrape', type: 'info' });
        await GetNewScrape(setScrapeData,setAlert)
      }
    } catch (error) {
      console.error(error);
      setAlert({ message: 'No existing scrape found trying to fetch new scrape', type: 'info' });
      await GetNewScrape(setScrapeData,setAlert)
    }
  };
  
  const GetNewScrape = async (setScrapeData , setAlert) => {
    try {
      setAlert({ message: 'New scrape Started ', type: 'success' });
      const response = await axios.post('http://localhost:8002/scrape');
      if(response.status ===200 && response.data.length>0){
        setAlert({ message: 'New Scrape Successful updated data and form', type: 'success' });
        setScrapeData(Parse(response.data))
        return
      }else {
        setAlert({ message: 'Scrape Failed Plz try again', type: 'warn' });
      }
    } catch (error) {
      setAlert({ message: 'Scrape Failed Plz try again', type: 'error' });
      console.error(error);
    }
  };

  const LoadNewScrape = (
    scrapeData,
    selectedCategory,
    selectedOperatingWt,
    selectedEngineHP,
    setMinOperatingWT,
    setMaxOperatingWT,
    setCategory,
    setMinEngineHP,
    setMaxEngineHP,
    setSelectedCategory,
    setSelectedOperatingWt,
    setSelectedEngineHP,
    setDisplayData
    ) => {
      if(selectedCategory === undefined )
      return

   var newCategory = [] 
   var minOperatingWT = 9999999999999; 
   var maxOperatingWT = 0;
   var minEngineHP = 999999999999999;
   var maxEngineHP = 0;

    for(var i in scrapeData) {
        var data = scrapeData[i];

        if(!newCategory.includes(data.category)){
          newCategory.push(data.category)
        }
        if(parseInt(data.weight) < minOperatingWT){
          minOperatingWT = data.weight
        }
        if(parseInt(data.weight) > maxOperatingWT){
          maxOperatingWT = data.weight
        }
        if(parseInt(data.hp) < minEngineHP){

          minEngineHP = data.hp
        }
        if(parseInt(data.hp) > maxEngineHP){
          maxEngineHP = data.hp
        }

        
    }
    setMinEngineHP(minEngineHP)
    setMaxEngineHP(maxEngineHP)
    setMinOperatingWT(minOperatingWT)
    setMaxOperatingWT(maxOperatingWT)

    setCategory(newCategory)
    setSelectedCategory(...selectedCategory.filter(c=> newCategory.includes(c) ))
   setSelectedOperatingWt(...[minOperatingWT<=selectedOperatingWt[0]<=maxOperatingWT?selectedOperatingWt[0]:minOperatingWT,
    minOperatingWT<=selectedOperatingWt[1]<=maxOperatingWT?selectedOperatingWt[1]:maxOperatingWT,
  ])

  setSelectedEngineHP(...[minEngineHP<=selectedEngineHP[0]<=maxEngineHP?selectedEngineHP[0]:minEngineHP,
    minEngineHP<=selectedEngineHP[1]<=maxEngineHP?selectedEngineHP[1]:maxEngineHP,
  ])
  setDisplayData([...UpdateFilter(scrapeData,selectedCategory,selectedOperatingWt,selectedEngineHP,setDisplayData)])
  }

  const UpdateFilter = (scrapeData,selectedCategory,selectedOperatingWt,selectedEngineHP,setDisplayData) => {
    var data =  scrapeData.filter(data => {
      return FilterData(data,selectedCategory,selectedOperatingWt,selectedEngineHP,setDisplayData)
    })
    return data
  }

  const FilterData = (data,selectedCategory,selectedOperatingWt,selectedEngineHP) => {
    console.log(selectedCategory)
    var hpRange = (parseInt(selectedEngineHP[0])<=parseInt(data.hp)) && ((parseInt(data.hp) <=parseInt(selectedEngineHP[1])))
    var opWt = (parseInt(selectedOperatingWt[0])<=parseInt(data.weight)) && ((parseInt(data.weight) <=parseInt(selectedOperatingWt[1])))
    if(selectedCategory.includes(data.category) && hpRange && opWt){
      return true
    }
    return false
  }

  

  export {GetNewScrape,GetScrape,LoadNewScrape,UpdateFilter}