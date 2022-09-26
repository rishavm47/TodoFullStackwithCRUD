// import './App.css';
import Main from './components/Main';

function App() {
  // const [todos, setTodos] = useState([]);
  // useEffect(() => {
  //   gettodos();
  //   // console.log("hello");
  // }, [])
  // async function gettodos() {
  //   try {
  //     const response = await axios.get('http://localhost:8080/todos');
  //     if (response.data)
  //       setTodos(response.data);
  //   } catch (error) {
  //     console.error(error);
  //   }
  // }

  return Main();
}

export default App;