import { useContext } from "react";
import { useNavigate } from "react-router-dom";
import { ModalContext } from "../../context/modalContext";
import GeneralModal from "../Modals/GeneralModal";
import CreateCourseForm from "./CreateCourseForm";

const CourseCard = () => {
    const {state, actions}=  useContext(ModalContext)
    const navigate = useNavigate()
    return (
        <>
            <GeneralModal>
                    <CreateCourseForm />
            </GeneralModal>

            <div className="card w-96 bg-base-200 shadow-xl mx-2 my-2">
                <div className="card-body">
                    <h2 className="card-title">Course</h2>
                    <p>Create a brand new course here or look up and edit any existing courses.</p>
                    <div className="card-actions justify-end">
                        <button className="btn btn-sm btn-primary"
                            onClick={()=>{
                                navigate("/course")
                            }}
                        >
                            Show All
                        </button>
                        <button onClick={()=>{
                            actions.openModal()
                        }} className="btn btn-sm btn-primary">New</button>
                    </div>
                </div>
            </div>
        </>

    );
}

export default CourseCard;