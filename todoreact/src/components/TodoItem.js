import React from 'react';
import { useDispatch } from 'react-redux';
import { deleteTodoAsync, updateTodoAsync } from '../redux/todoSlice';
import '../styles/TodoItem.css'
// import { toggleComplete } from '../redux/todoSlice';

const TodoItem = ({ getColor, id, title, completed, handleEditClickMain }) => {
    // console.log(completed);
    const dispatch = useDispatch();
    // const [comp, setComp] = useState(completed);

    const handleDeleteClick = () => {
        // console.log("here");
        dispatch(deleteTodoAsync({ id }));
    };
    const handleToggleComp = () => {
        dispatch(updateTodoAsync({
            id, title, completed: !completed
        }))
        // setComp(!completed);
    }

    return (
        <li className='Applist' key={id} >
            <input style={{ marginRight: '20px' }} checked={completed} type='checkbox' onChange={handleToggleComp}></input>
            <div style={{ width: 180, borderWidth: 1, borderColor: 'red', overflow: "scroll"}}>{title}</div>

            <button onClick={() => handleEditClickMain({ title, id })}>Edit</button>
            <button onClick={handleDeleteClick}>X</button>
        </li>
    );
};

export default TodoItem;