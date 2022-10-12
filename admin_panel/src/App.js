import Home from './pages/Home'
import {
  createBrowserRouter,
  RouterProvider,
  Route,
} from "react-router-dom";
import CourseForm from './components/Course/CourseForm';
import Course from './pages/Course';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Course/>,
  },
]);

function App() {

  return (
    <>
      <RouterProvider router={router} />
    </>

  );
}

export default App;
