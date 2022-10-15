import { useNavigate } from 'react-router-dom';
import { FaTrashAlt } from 'react-icons/fa';
import { useContext } from 'react';
import { ConfirmModalCtx } from '../../context/confirmModalCtx';
import ConfirmationModal from '../Modals/ConfirmationModal';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { deleteCourse } from '../../services/api_service';


const CourseListCard = ({ course }) => {
    const navigate = useNavigate()
    const confirmModalCtx = useContext(ConfirmModalCtx)
    const queryClient = useQueryClient()

    const mutation = useMutation(deleteCourse, {
        onError: (error, variables, context) => {
            // An error happened!
            console.log(error)
        },
        onSuccess: ({ data }, variables, context) => {
            queryClient.invalidateQueries(["getAllCourses"])
        },
    })


    return (
        <>
            <ConfirmationModal/>
            <div className="card w-80 bg-base-100 shadow-xl" key={course.ID}>
                <figure><img src="https://placeimg.com/400/225/arch" alt="Shoes" /></figure>
                <div className="card-body">
                    <h3 className="card-title cursor-pointer hover:underline"
                        onClick={
                            () => {
                                navigate(`/course/${course.ID}`)
                            }
                        }
                    >
                        {course.Title}
                    </h3>
                    <p className="text-sm">{course.DescShort}</p>

                    <div className="card-actions justify-end">
                        <div className="badge badge-outline">Backend</div>
                        <div className="badge badge-outline">Software Development</div>
                    </div>
                    <a className="btn btn-sm self-end"
                        onClick={()=>{
                            confirmModalCtx.actions.openModal()
                            confirmModalCtx.actions.setOnYesAction({
                                action: ()=>{
                                    mutation.mutate(course.ID, "delete-course")
                                }
                            })
                        }}
                    >
                        <FaTrashAlt />
                    </a>
                </div>
            </div>
        </>
    );
}

export default CourseListCard;