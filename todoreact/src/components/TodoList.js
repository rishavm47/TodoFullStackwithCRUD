import React, { useEffect } from 'react';
import TodoItem from './TodoItem';
import { useDispatch, useSelector } from 'react-redux';
import { getTodosAsync } from '../redux/todoSlice';

const TodoList = ({ handleEditClickMain }) => {
    const todos = useSelector((state) => state.todos);
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(getTodosAsync());
        // console.log(todos.completed);
        // console.log(todos[2].);
    }, [dispatch])



    return (
        <ul className='list-group'>
            {todos.slice().reverse().map((todo) => (
                <TodoItem id={todo.id} title={todo.title} completed={todo.completed} handleEditClickMain={handleEditClickMain} />
            ))}
        </ul>
    );
};

export default TodoList;