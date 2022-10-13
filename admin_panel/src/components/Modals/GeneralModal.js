const GeneralModal = ({open, children}) => {
    return (
        <div className={`modal ${open?"modal-open":""}`} >
            <div className="modal-box relative">
                <a href="#" className="btn btn-sm btn-circle absolute right-2 top-2">✕</a>
                {children}
            </div>
        </div>
    );
}

export default GeneralModal;