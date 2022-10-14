import Home from './pages/Home'
import {
  createBrowserRouter,
  RouterProvider,
  Route,
} from "react-router-dom";
import {
  useQuery,
  useMutation,
  useQueryClient,
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'

import CourseForm from './components/Course/CourseForm';
import Course from './pages/Course';
import CourseList from './pages/CourseList';
import ModalContextProvider from './context/modalContext';
import ConfirmModalCtxProvider, { ConfirmModalCtx } from './context/confirmModalCtx';

const queryClient = new QueryClient()

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/course/:course_id",
    element: <Course />,
  },
  {
    path: "/course/",
    element: <CourseList />,
  },
  {
    path: "/course/:course_id/post/:post_id",
    element: <Course />,
  },
]);

function App() {

  return (
    <>
      <QueryClientProvider client={queryClient}>
        <ReactQueryDevtools initialIsOpen={true} />
        <ModalContextProvider>
          <ConfirmModalCtxProvider>
            <RouterProvider router={router} />
          </ConfirmModalCtxProvider>
        </ModalContextProvider>
      </QueryClientProvider>
    </>
  );
}

export default App;
