import { useState } from 'react';
import { useDispatch } from "react-redux";
import '../App.css';
import { addTodoAsync, updateTodoAsync } from "../redux/todoSlice";
import TodoList from "./TodoList";


const Main = () => {



    const [isEditing, setIsEditing] = useState(false);
    const [currentTodo, setCurrentTodo] = useState({});

    const [field, setField] = useState("");
    const dispatch = useDispatch();



    function handleEditInputChange(e) {
        setCurrentTodo({ ...currentTodo, title: e.target.value });
        // console.log(currentTodo);
    }

    function handleInputChange(e) {
        setField(e.target.value);
    }

    function handleEditClickMain(todo) {
        setIsEditing(true);
        setCurrentTodo({ ...todo });
    }


    function handleFormSubmit(e) {
        e.preventDefault();

        if (field) {
            dispatch(
                addTodoAsync({
                    title: field,
                })
            );

        }
        setField("");

    }
    function handleUpdateTodo(id, updatedTodo) {
        // const updatedItem = todos.map((todo) => {
        //     if (todo.id === id) {
        //         axios.patch(`http://localhost:8080/todos/${id}`, {
        //             title: updatedTodo.title
        //         })
        //         return updatedTodo
        //     }
        //     else
        //         return todo;
        // });
        dispatch(updateTodoAsync(updatedTodo));
        setIsEditing(false);
        // setTodos(updatedItem);
    }
    function handleEditFormSubmit(e) {
        e.preventDefault();

        handleUpdateTodo(currentTodo.id, currentTodo);
    }

    return (
        <div className="App">
            <header className="App-header">
                {
                    isEditing ? (
                        <>
                            <h2 style={{ color: '#c591f9' }}>Edit Todo</h2>
                            <form style={{ display: "flex" }} onSubmit={handleEditFormSubmit}>
                                <input
                                    name="editTodo"
                                    type="text"
                                    placeholder="Edit todo"
                                    value={currentTodo.title}
                                    onChange={handleEditInputChange}
                                />
                                <button type="submit">Update</button>
                                <button onClick={() => setIsEditing(false)}>Cancel</button>
                            </form>
                        </>
                    ) : (
                        <>
                            <h1 className="App-logo">todos</h1>
                            <div>
                                <form onSubmit={handleFormSubmit}>
                                    <input onChange={handleInputChange} value={field} type='text' placeholder='Enter the task' />

                                </form>
                                <TodoList handleEditClickMain={handleEditClickMain} />
                            </div>
                        </>
                    )
                }

            </header>
        </div>
    );
}

export default Main;