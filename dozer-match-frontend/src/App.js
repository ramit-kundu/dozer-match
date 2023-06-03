import logo from './logo.svg';
import './App.css';
import Dozer  from './components/Dozer/Dozer';
import DozerList from './containers/DozerList/DozerList';
import Home from './containers/Home/Home';

function App() {

  // const category = 'Excavator';
  // const engine = 'Diesel';
  // const hp = '200';
  // const weight = '5000kg';
  // const make = 'ABC Company';
  // const model = 'XYZ Model';
  // const imageUrl = 'https://s7d2.scene7.com/is/image/Caterpillar/CM20200429-439d6-30cb1';
  return (
    <div className="App">
      <Home/>
    </div>
  );
}

export default App;
