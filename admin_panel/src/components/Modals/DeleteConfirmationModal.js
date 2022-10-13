const DeleteConfirmationModal = () => {
    return (
        <div className="modal" id="delete-confirmation-modal">
            <div className="modal-box">
                <h3 className="font-bold text-lg">Are you sure you want to delete?</h3>
                <div className="modal-action">
                    <a href="#" className="btn">Yes</a>
                    <a href="#" className="btn">No</a>
                </div>
            </div>
        </div>
    );
}

export default DeleteConfirmationModal;