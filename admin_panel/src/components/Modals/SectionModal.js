import { useContext } from "react";
import { SectionContext } from "../../context/sectionContext";
import SectionForm from "../Course/SectionForm";

const SectionModal = ({children}) => {
    const {state, actions}= useContext(SectionContext)
    return (
        <div className={`modal ${state.isOpen?"modal-open":""}`} >
            <div className="modal-box relative">
                <a className="btn btn-sm btn-circle absolute right-2 top-2"
                    onClick={()=>{
                        actions.closeModal()
                    }}
                >
                    ✕
                </a>
                <SectionForm/>
            </div>
        </div>
    );
}

export default SectionModal;