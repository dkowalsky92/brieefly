-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2018-10-21 23:42:26.656

-- tables
-- Table: Agency
CREATE TABLE Agency (
    id_company varchar(100) NOT NULL,
    agency_code varchar(13) NOT NULL,
    nip_number int NULL,
    CONSTRAINT Agency_pk PRIMARY KEY (id_company)
);

-- Table: Agency_employee
CREATE TABLE Agency_employee (
    id_user varchar(100) NOT NULL,
    position varchar(70) NULL,
    id_company varchar(100) NOT NULL,
    CONSTRAINT Agency_employee_pk PRIMARY KEY (id_user)
);

-- Table: Agency_employee_project_phase
CREATE TABLE Agency_employee_project_phase (
    id_agency_employee_project_phase varchar(100) NOT NULL,
    id_user varchar(100) NOT NULL,
    id_phase varchar(100) NOT NULL,
    date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT Agency_employee_project_phase_pk PRIMARY KEY (id_agency_employee_project_phase)
);

-- Table: Answer_option
CREATE TABLE Answer_option (
    id_answer_option varchar(100) NOT NULL,
    content varchar(150) NOT NULL,
    is_chosen bool NOT NULL DEFAULT false,
    id_question varchar(100) NOT NULL,
    CONSTRAINT Answer_option_pk PRIMARY KEY (id_answer_option)
);

-- Table: Client
CREATE TABLE Client (
    id_user varchar(100) NOT NULL,
    job varchar(40) NULL,
    id_company varchar(100) NULL,
    CONSTRAINT Client_pk PRIMARY KEY (id_user)
);

-- Table: Client_agency
CREATE TABLE Client_agency (
    id_client_agency varchar(100) NOT NULL,
    id_user varchar(100) NOT NULL,
    id_company varchar(100) NOT NULL,
    date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT Client_agency_pk PRIMARY KEY (id_client_agency)
);

-- Table: Client_project
CREATE TABLE Client_project (
    id_client_project varchar(100) NOT NULL,
    id_user varchar(100) NOT NULL,
    id_project varchar(100) NOT NULL,
    date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT Client_project_pk PRIMARY KEY (id_client_project)
);

-- Table: Cms
CREATE TABLE Cms (
    id_cms varchar(100) NOT NULL,
    name varchar(40) NOT NULL,
    description varchar(300) NULL,
    CONSTRAINT Cms_pk PRIMARY KEY (id_cms)
);

-- Table: Color
CREATE TABLE Color (
    id_color varchar(100) NOT NULL,
    hex_value varchar(8) NOT NULL,
    id_project varchar(100) NULL,
    CONSTRAINT Color_pk PRIMARY KEY (id_color)
);

-- Table: Company
CREATE TABLE Company (
    id_company varchar(100) NOT NULL,
    name varchar(100) NOT NULL,
    url_name varchar(100) NOT NULL UNIQUE,
    website_url varchar(300) NULL,
    phone varchar(20) NULL,
    email varchar(75) NOT NULL,
    address varchar(300) NULL,
    description varchar(300) NULL,
    image_url varchar(200) NULL,
    date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_last_modified timestamp NULL ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT Company_pk PRIMARY KEY (id_company)
);

-- Table: Custom_feature
CREATE TABLE Custom_feature (
    id_custom_feature varchar(100) NOT NULL,
    name varchar(50) NOT NULL,
    description varchar(500) NOT NULL,
    id_project varchar(100) NULL,
    CONSTRAINT Custom_feature_pk PRIMARY KEY (id_custom_feature)
);

-- Table: Feature
CREATE TABLE Feature (
    id_feature varchar(100) NOT NULL,
    name varchar(50) NOT NULL,
    description varchar(300) NULL,
    CONSTRAINT Feature_pk PRIMARY KEY (id_feature)
);

-- Table: Offer
CREATE TABLE Offer (
    id_offer varchar(100) NOT NULL,
    salary int NOT NULL,
    is_chosen bool NOT NULL DEFAULT false,
    date_deadline date NOT NULL,
    date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    id_project varchar(100) NULL,
    id_company varchar(100) NOT NULL,
    CONSTRAINT Offer_pk PRIMARY KEY (id_offer)
);

-- Table: Opinion
CREATE TABLE Opinion (
    id_opinion varchar(100) NOT NULL,
    grade int NOT NULL,
    description varchar(500) NULL,
    date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    id_project varchar(100) NOT NULL,
    CONSTRAINT Opinion_pk PRIMARY KEY (id_opinion)
);

-- Table: Project
CREATE TABLE Project (
    id_project varchar(100) NOT NULL,
    name varchar(70) NOT NULL,
    url_name varchar(70) NOT NULL UNIQUE,
    type varchar(40) NOT NULL,
    description varchar(500) NOT NULL,
    language varchar(50) NULL,
    budget_min int NULL,
    budget_max int NULL,
    subpage_count int NULL,
    overall_progress int NOT NULL DEFAULT 0,
    image_url varchar(100) NULL,
    date_deadline date NULL,
    date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_last_modified timestamp NULL ON UPDATE CURRENT_TIMESTAMP,
    id_status varchar(100) NOT NULL,
    id_cms varchar(100) NULL,
    CONSTRAINT Project_pk PRIMARY KEY (id_project)
);

-- Table: Project_feature
CREATE TABLE Project_feature (
    id_project varchar(100) NOT NULL,
    id_feature varchar(100) NOT NULL,
    CONSTRAINT Project_feature_pk PRIMARY KEY (id_project,id_feature)
);

-- Table: Phase
CREATE TABLE Phase (
    id_phase varchar(100) NOT NULL,
    name varchar(40) NOT NULL,
    description varchar(300) NULL,
    value int NOT NULL DEFAULT 1,
    progress int NOT NULL DEFAULT 0,
    order_position int NULL,
    status varchar(20) NOT NULL,
    date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    id_phase varchar(100) NOT NULL,
    id_project varchar(100) NOT NULL,
    CONSTRAINT Phase_pk PRIMARY KEY (id_phase)
);

-- Table: Question
CREATE TABLE Question (
    id_question varchar(100) NOT NULL,
    type varchar(100) NOT NULL,
    content varchar(500) NOT NULL,
    status varchar(30) NOT NULL DEFAULT 'pending',
    date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_last_modified timestamp NULL ON UPDATE CURRENT_TIMESTAMP,
    id_user varchar(100) NOT NULL,
    id_phase varchar(100) NOT NULL,
    CONSTRAINT Question_pk PRIMARY KEY (id_question)
);

-- Table: Similar_project
CREATE TABLE Similar_project (
    id_similar_project varchar(100) NOT NULL,
    project_url varchar(300) NOT NULL,
    id_project varchar(100) NULL,
    CONSTRAINT Similar_project_pk PRIMARY KEY (id_similar_project)
);

-- Table: Status
CREATE TABLE Status (
    id_status varchar(100) NOT NULL,
    name varchar(30) NOT NULL,
    CONSTRAINT Status_pk PRIMARY KEY (id_status)
);

-- Table: Target_group
CREATE TABLE Target_group (
    id_target_group varchar(100) NOT NULL,
    name varchar(50) NOT NULL,
    description varchar(500) NULL,
    age_min int NULL,
    age_max int NULL,
    id_project varchar(100) NULL,
    CONSTRAINT Target_group_pk PRIMARY KEY (id_target_group)
);

-- Table: Task
CREATE TABLE Task (
    id_task varchar(100) NOT NULL,
    name varchar(30) NOT NULL,
    value int NOT NULL DEFAULT 1,
    is_done bool NOT NULL DEFAULT false,
    id_phase varchar(100) NOT NULL,
    CONSTRAINT Task_pk PRIMARY KEY (id_task)
);

-- Table: User
CREATE TABLE User (
    id_user varchar(100) NOT NULL,
    login varchar(20) NULL,
    password varchar(20) NULL,
    password_fail_attempts int NOT NULL DEFAULT 0,
    email varchar(75) NOT NULL UNIQUE,
    name varchar(20) NULL,
    surname varchar(20) NULL,
    phone varchar(14) NULL,
    website_url varchar(300) NULL,
    image_url varchar(200) NULL,
    description varchar(300) NULL,
    date_of_birth date NULL,
    date_last_logged date NULL,
    date_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_last_modified timestamp NULL ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT User_pk PRIMARY KEY (id_user)
);

-- Table: Visual_identity
CREATE TABLE Visual_identity (
    id_visual_identity varchar(100) NOT NULL,
    type varchar(30) NOT NULL,
    id_project varchar(100) NULL,
    CONSTRAINT Visual_identity_pk PRIMARY KEY (id_visual_identity)
);

-- foreign keys
-- Reference: Agency_Company (table: Agency)
ALTER TABLE Agency ADD CONSTRAINT Agency_Company FOREIGN KEY Agency_Company (id_company)
    REFERENCES Company (id_company)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Agency_employee_Agency (table: Agency_employee)
ALTER TABLE Agency_employee ADD CONSTRAINT Agency_employee_Agency FOREIGN KEY Agency_employee_Agency (id_company)
    REFERENCES Agency (id_company)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Agency_employee_User (table: Agency_employee)
ALTER TABLE Agency_employee ADD CONSTRAINT Agency_employee_User FOREIGN KEY Agency_employee_User (id_user)
    REFERENCES User (id_user)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Agency_employee_project_phase_Agency_employee (table: Agency_employee_project_phase)
ALTER TABLE Agency_employee_project_phase ADD CONSTRAINT Agency_employee_project_phase_Agency_employee FOREIGN KEY Agency_employee_project_phase_Agency_employee (id_user)
    REFERENCES Agency_employee (id_user)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Agency_employee_project_phase_Project_phase (table: Agency_employee_project_phase)
ALTER TABLE Agency_employee_project_phase ADD CONSTRAINT Agency_employee_project_phase_Phase FOREIGN KEY Agency_employee_project_phase_Project_phase (id_phase)
    REFERENCES Project_phase (id_phase)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Client_Company (table: Client)
ALTER TABLE Client ADD CONSTRAINT Client_Company FOREIGN KEY Client_Company (id_company)
    REFERENCES Company (id_company)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Client_User (table: Client)
ALTER TABLE Client ADD CONSTRAINT Client_User FOREIGN KEY Client_User (id_user)
    REFERENCES User (id_user)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Client_agency_Agency (table: Client_agency)
ALTER TABLE Client_agency ADD CONSTRAINT Client_agency_Agency FOREIGN KEY Client_agency_Agency (id_company)
    REFERENCES Agency (id_company)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Client_agency_Client (table: Client_agency)
ALTER TABLE Client_agency ADD CONSTRAINT Client_agency_Client FOREIGN KEY Client_agency_Client (id_user)
    REFERENCES Client (id_user)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Client_project_Client (table: Client_project)
ALTER TABLE Client_project ADD CONSTRAINT Client_project_Client FOREIGN KEY Client_project_Client (id_user)
    REFERENCES Client (id_user)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Client_project_Project (table: Client_project)
ALTER TABLE Client_project ADD CONSTRAINT Client_project_Project FOREIGN KEY Client_project_Project (id_project)
    REFERENCES Project (id_project)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Color_Project (table: Color)
ALTER TABLE Color ADD CONSTRAINT Color_Project FOREIGN KEY Color_Project (id_project)
    REFERENCES Project (id_project)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Custom_feature_Project (table: Custom_feature)
ALTER TABLE Custom_feature ADD CONSTRAINT Custom_feature_Project FOREIGN KEY Custom_feature_Project (id_project)
    REFERENCES Project (id_project)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Offer_Agency (table: Offer)
ALTER TABLE Offer ADD CONSTRAINT Offer_Agency FOREIGN KEY Offer_Agency (id_company)
    REFERENCES Agency (id_company)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Offer_Project (table: Offer)
ALTER TABLE Offer ADD CONSTRAINT Offer_Project FOREIGN KEY Offer_Project (id_project)
    REFERENCES Project (id_project)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Opinion_Project (table: Opinion)
ALTER TABLE Opinion ADD CONSTRAINT Opinion_Project FOREIGN KEY Opinion_Project (id_project)
    REFERENCES Project (id_project)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Option_Question (table: Answer_option)
ALTER TABLE Answer_option ADD CONSTRAINT Option_Question FOREIGN KEY Option_Question (id_question)
    REFERENCES Question (id_question);

-- Reference: Project_Cms (table: Project)
ALTER TABLE Project ADD CONSTRAINT Project_Cms FOREIGN KEY Project_Cms (id_cms)
    REFERENCES Cms (id_cms)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Project_Status (table: Project)
ALTER TABLE Project ADD CONSTRAINT Project_Status FOREIGN KEY Project_Status (id_status)
    REFERENCES Status (id_status)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Project_feature_Feature (table: Project_feature)
ALTER TABLE Project_feature ADD CONSTRAINT Project_feature_Feature FOREIGN KEY Project_feature_Feature (id_feature)
    REFERENCES Feature (id_feature)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Project_feature_Project (table: Project_feature)
ALTER TABLE Project_feature ADD CONSTRAINT Project_feature_Project FOREIGN KEY Project_feature_Project (id_project)
    REFERENCES Project (id_project)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Project_phase_Project (table: Project_phase)
ALTER TABLE Phase ADD CONSTRAINT Phase_Project FOREIGN KEY Phase_Project (id_project)
    REFERENCES Project (id_project)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Question_Agency_employee (table: Question)
ALTER TABLE Question ADD CONSTRAINT Question_Agency_employee FOREIGN KEY Question_Agency_employee (id_user)
    REFERENCES Agency_employee (id_user)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Question_Client (table: Question)
ALTER TABLE Question ADD CONSTRAINT Question_Client FOREIGN KEY Question_Client (id_user)
    REFERENCES Client (id_user)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Question_Project_phase (table: Question)
ALTER TABLE Question ADD CONSTRAINT Question_Phase FOREIGN KEY Question_Phase (id_phase)
    REFERENCES Phase (id_phase)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Similar_project_Project (table: Similar_project)
ALTER TABLE Similar_project ADD CONSTRAINT Similar_project_Project FOREIGN KEY Similar_project_Project (id_project)
    REFERENCES Project (id_project)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Target_group_Project (table: Target_group)
ALTER TABLE Target_group ADD CONSTRAINT Target_group_Project FOREIGN KEY Target_group_Project (id_project)
    REFERENCES Project (id_project)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Task_Project_phase (table: Task)
ALTER TABLE Task ADD CONSTRAINT Task_Phase FOREIGN KEY Task_Phase (id_phase)
    REFERENCES Phase (id_phase)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- Reference: Visual_identity_Project (table: Visual_identity)
ALTER TABLE Visual_identity ADD CONSTRAINT Visual_identity_Project FOREIGN KEY Visual_identity_Project (id_project)
    REFERENCES Project (id_project)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- -- Reference: Agency_employee_Agency_employee (table: Agency_employee)
-- ALTER TABLE Agency_employee ADD CONSTRAINT Agency_employee_Agency_employee FOREIGN KEY Agency_employee_Agency_employee (id_user)
--     REFERENCES Agency_employee (id_user)
--     ON DELETE CASCADE
--     ON UPDATE CASCADE;

-- End of file.

